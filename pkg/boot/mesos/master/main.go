package main

import (
	"fmt"
	"strings"

	"github.com/aledbf/coreos-mesos-zookeeper/pkg/boot"
	"github.com/aledbf/coreos-mesos-zookeeper/pkg/etcd"
	logger "github.com/aledbf/coreos-mesos-zookeeper/pkg/log"
	"github.com/aledbf/coreos-mesos-zookeeper/pkg/os"
	"github.com/aledbf/coreos-mesos-zookeeper/pkg/types"
	goetcd "github.com/coreos/go-etcd/etcd"
)

const (
	mesosPort = 5050
)

var (
	etcdPath = os.Getopt("ETCD_PATH", "/deis/mesos/master")
	log      = logger.New()
)

func init() {
	boot.RegisterComponent(new(MesosBoot), "boot")
}

func main() {
	boot.Start(etcdPath, mesosPort)
}

type MesosBoot struct{}

func (mb *MesosBoot) MkdirsEtcd() []string {
	return []string{etcdPath}
}

func (mb *MesosBoot) EtcdDefaults() map[string]string {
	return map[string]string{}
}

func (mb *MesosBoot) PreBootScripts(currentBoot *types.CurrentBoot) []*types.Script {
	return []*types.Script{}
}

func (mb *MesosBoot) PreBoot(currentBoot *types.CurrentBoot) {
	log.Info("mesos-master: starting...")
}

func (mb *MesosBoot) BootDaemons(currentBoot *types.CurrentBoot) []*types.ServiceDaemon {
	return []*types.ServiceDaemon{&types.ServiceDaemon{Command: "mesos-master", Args: gatherArgs(currentBoot.EtcdClient)}}
}

func (mb *MesosBoot) WaitForPorts() []int {
	return []int{mesosPort}
}

func (mb *MesosBoot) PostBootScripts(currentBoot *types.CurrentBoot) []*types.Script {
	return []*types.Script{}
}

func (mb *MesosBoot) PostBoot(currentBoot *types.CurrentBoot) {
	log.Info("mesos-master: running...")
}

func (mb *MesosBoot) ScheduleTasks(currentBoot *types.CurrentBoot) []*types.Cron {
	return []*types.Cron{}
}

func (mb *MesosBoot) UseConfd() bool {
	return false
}

func (mb *MesosBoot) PreShutdownScripts(currentBoot *types.CurrentBoot) []*types.Script {
	return []*types.Script{}
}

func gatherArgs(c *goetcd.Client) []string {
	var args []string

	nodes := etcd.GetList(c, "/zookeeper/nodes")
	var hosts []string
	for _, node := range nodes {
		hosts = append(hosts, node+":3888")
	}
	zkHosts := strings.Join(hosts, ",")
	args = append(args, "--zk="+zkHosts+"/mesos")

	// set quorum based on num zk hosts
	l := len(nodes)
	args = append(args, fmt.Sprintf("--quorum=%v", l/2+1))
	// set a work directory
	args = append(args, "--work_dir=/tmp/mesos-master")

	return args
}
