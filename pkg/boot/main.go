//go:generate go-extpoints
package boot

import (
	"net"
	"net/http"
	"os"
	"os/signal"
	"runtime"
	"strconv"
	"strings"
	"syscall"
	"time"

	"github.com/aledbf/coreos-mesos-zookeeper/pkg/boot/extpoints"

	"github.com/aledbf/coreos-mesos-zookeeper/pkg/confd"
	"github.com/aledbf/coreos-mesos-zookeeper/pkg/etcd"
	logger "github.com/aledbf/coreos-mesos-zookeeper/pkg/log"
	netwrapper "github.com/aledbf/coreos-mesos-zookeeper/pkg/net"
	oswrapper "github.com/aledbf/coreos-mesos-zookeeper/pkg/os"
	"github.com/aledbf/coreos-mesos-zookeeper/pkg/types"
	"github.com/robfig/cron"
	_ "net/http/pprof"
)

const (
	timeout  time.Duration = 10 * time.Second
	ttl      time.Duration = timeout * 2
	etcdPort int           = 4001
)

var (
	signalChan  = make(chan os.Signal, 1)
	log         = logger.New()
	bootProcess = extpoints.BootComponents
	component   extpoints.BootComponent
)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

// RegisterComponent register an externsion to be used with this application
func RegisterComponent(component extpoints.BootComponent, name string) bool {
	return bootProcess.Register(component, name)
}

// Start initiate the boot process of the current component
// etcdPath is the base path used to publish the component in etcd
// externalPort is the base path used to publish the component in etcd
func Start(etcdPath string, externalPort int) {
	go func() {
		log.Debugf("starting pprof http server in port 6060")
		http.ListenAndServe("localhost:6060", nil)
	}()

	signal.Notify(signalChan,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGKILL,
		syscall.SIGTERM,
		syscall.SIGQUIT,
		os.Interrupt,
	)

	// Wait for a signal and exit
	exitChan := make(chan int)
	go func() {
		for {
			s := <-signalChan
			log.Debugf("Signal received: %v", s)
			switch s {
			case syscall.SIGTERM:
				exitChan <- 0
			case syscall.SIGQUIT:
				exitChan <- 0
			case syscall.SIGKILL:
				exitChan <- 1
			default:
				exitChan <- 1
			}
		}
	}()

	component = bootProcess.Lookup("boot")
	if component == nil {
		log.Error("error loading boot extension...")
		signalChan <- syscall.SIGINT
	}

	host := oswrapper.Getopt("HOST", "127.0.0.1")
	etcdCtlPeers := oswrapper.Getopt("ETCDCTL_PEERS", "127.0.0.1")
	etcdClient := etcd.NewClient(getHttpEtcdUrls(host, etcdCtlPeers, etcdPort))

	currentBoot := &types.CurrentBoot{
		ConfdNodes: getConfdNodes(host, etcdCtlPeers, etcdPort),
		EtcdClient: etcdClient,
		EtcdPath:   etcdPath,
		EtcdPort:   etcdPort,
		Host:       net.ParseIP(host),
		Timeout:    timeout,
		TTL:        timeout * 2,
		Port:       externalPort,
	}

	// do the real work in a goroutine to be able to exit if
	// a signal is received during the boot process
	go start(currentBoot)

	code := <-exitChan

	// pre shutdown tasks
	log.Debugf("executing pre shutdown scripts")
	preShutdownScripts := component.PreShutdownScripts(currentBoot)
	runAllScripts(signalChan, preShutdownScripts)

	log.Debugf("execution terminated with exit code %v", code)
	os.Exit(code)
}

func start(currentBoot *types.CurrentBoot) {
	log.Info("starting component...")

	for _, key := range component.MkdirsEtcd() {
		etcd.Mkdir(currentBoot.EtcdClient, key)
	}

	for key, value := range component.EtcdDefaults() {
		etcd.SetDefault(currentBoot.EtcdClient, key, value)
	}

	component.PreBoot(currentBoot)

	if component.UseConfd() {
		// wait until etcd has discarded potentially stale values
		time.Sleep(timeout + 1)

		// wait for confd to run once and install initial templates
		confd.WaitForInitialConf(currentBoot.ConfdNodes, currentBoot.Timeout)
	}

	log.Debug("running pre boot scripts")
	preBootScripts := component.PreBootScripts(currentBoot)
	runAllScripts(signalChan, preBootScripts)

	if component.UseConfd() {
		// spawn confd in the background to update services based on etcd changes
		go confd.Launch(signalChan, currentBoot.ConfdNodes)
	}

	log.Debug("running boot daemons")
	servicesToStart := component.BootDaemons(currentBoot)
	for _, daemon := range servicesToStart {
		go oswrapper.RunProcessAsDaemon(signalChan, daemon.Command, daemon.Args)
	}

	portsToWaitFor := component.WaitForPorts()
	log.Debugf("waiting for a service in the port %v", portsToWaitFor)
	for _, portToWait := range portsToWaitFor {
		if portToWait > 0 {
			err := netwrapper.WaitForPort("tcp", "0.0.0.0", strconv.Itoa(portToWait), timeout)
			if err != nil {
				log.Printf("%v", err)
				signalChan <- syscall.SIGINT
			}
		}
	}

	// we only publish the service in etcd if the port if > 0
	if currentBoot.Port > 0 {
		log.Debug("starting periodic publication in etcd...")
		log.Debugf("etcd publication path %s, host %s and port %v", currentBoot.EtcdPath, currentBoot.Host, currentBoot.Port)
		go etcd.PublishService(currentBoot.EtcdClient, currentBoot.Host.String(), currentBoot.EtcdPath, currentBoot.Port, uint64(ttl.Seconds()), timeout)
	}

	// Wait for the first publication
	time.Sleep(timeout / 2)

	log.Printf("running post boot scripts")
	postBootScripts := component.PostBootScripts(currentBoot)
	runAllScripts(signalChan, postBootScripts)

	log.Debug("checking for cron tasks...")
	crons := component.ScheduleTasks(currentBoot)
	_cron := cron.New()
	for _, cronTask := range crons {
		_cron.AddFunc(cronTask.Frequency, cronTask.Code)
	}
	_cron.Start()

	component.PostBoot(currentBoot)
}

// getEtcdHosts returns an array of urls that contains at least one host
func getHttpEtcdUrls(host, etcdCtlPeers string, port int) []string {
	if etcdCtlPeers != "127.0.0.1" {
		hosts := strings.Split(etcdCtlPeers, ",")
		result := []string{}
		for _, _host := range hosts {
			result = append(result, "http://"+_host+":"+strconv.Itoa(port))
		}
		return result
	} else {
		return []string{"http://" + host + ":" + strconv.Itoa(port)}
	}
}

func getConfdNodes(host, etcdCtlPeers string, port int) []string {
	if etcdCtlPeers != "127.0.0.1" {
		hosts := strings.Split(etcdCtlPeers, ",")
		result := []string{}
		for _, _host := range hosts {
			result = append(result, _host+":"+strconv.Itoa(port))
		}
		return result
	} else {
		return []string{host + ":" + strconv.Itoa(port)}
	}
}

func runAllScripts(signalChan chan os.Signal, scripts []*types.Script) {
	for _, script := range scripts {
		if script.Params == nil {
			script.Params = map[string]string{}
		}
		// add HOME variable to avoid warning from ceph commands
		script.Params["HOME"] = "/tmp"
		if log.Level.String() == "debug" {
			script.Params["DEBUG"] = "true"
		}
		err := oswrapper.RunScript(script.Name, script.Params, script.Content)
		if err != nil {
			log.Printf("command finished with error: %v", err)
			signalChan <- syscall.SIGTERM
		}
	}
}
