package main

import (
	"strconv"

	"github.com/Scalingo/go-etcd-lock/lock"
	"github.com/aledbf/coreos-mesos-zookeeper/pkg/boot"
	"github.com/aledbf/coreos-mesos-zookeeper/pkg/etcd"
	"github.com/aledbf/coreos-mesos-zookeeper/pkg/fleet"
	logger "github.com/aledbf/coreos-mesos-zookeeper/pkg/log"
	"github.com/aledbf/coreos-mesos-zookeeper/pkg/os"
	"github.com/aledbf/coreos-mesos-zookeeper/pkg/types"
	goetcd "github.com/coreos/go-etcd/etcd"
)

const (
	fleetEndpoint = "unix:///var/run/fleet.sock"
)

var (
	etcdPath = os.Getopt("ETCD_PATH", "/zookeeper/nodes")
	log      = logger.New()
)

func init() {
	boot.RegisterComponent(new(ZkBoot), "boot")
}

func main() {
	boot.Start(etcdPath, 2181)
}

type ZkBoot struct{}

func (cb *ZkBoot) MkdirsEtcd() []string {
	return []string{etcdPath}
}

func (cb *ZkBoot) EtcdDefaults() map[string]string {
	return map[string]string{}
}

func (cb *ZkBoot) PreBootScripts(currentBoot *types.CurrentBoot) []*types.Script {
	return []*types.Script{}
}

func (cb *ZkBoot) PreBoot(currentBoot *types.CurrentBoot) {
	log.Info("zookeeper: starting...")
	// check if the nodes with the required role already have the an id. If not
	// get fleet nodes with the required role and preassing the ids for every node
	l, err := lock.Acquire(currentBoot.EtcdClient, "/zookeeper/masterLock", 120)
	if lockErr, ok := err.(*lock.Error); ok {
		log.Debug(lockErr)
		return
	} else if err != nil {
		panic(err)
	}

	zkNodes := etcd.GetList(currentBoot.EtcdClient, etcdPath)

	log.Debug("initializing zookeeper cluster ids...")
	machines, err := getMachines()
	if err != nil {
		panic(err)
	}

	if len(zkNodes) == 0 {
		// initialize cluster
		for index, newZkNode := range machines {
			etcd.Set(currentBoot.EtcdClient, etcdPath+"/"+newZkNode+"/id", strconv.Itoa(index+1), 0)
		}
	} else {
		// we check if some machine in the fleet cluster with the
		// required role is not initialized (no zookeeper node id).
		machinesNotInitialized := difference(machines, zkNodes)
		if len(machinesNotInitialized) > 0 {
			nextNodeId := getNextNodeId(currentBoot.EtcdClient, zkNodes)
			for _, zkNode := range machinesNotInitialized {
				etcd.Set(currentBoot.EtcdClient, etcdPath+"/"+zkNode+"/id", strconv.Itoa(nextNodeId), 0)
				nextNodeId++
			}
		}
	}

	l.Release()
}

func (cb *ZkBoot) BootDaemons(currentBoot *types.CurrentBoot) []*types.ServiceDaemon {
	cmd, args := os.BuildCommandFromString("/opt/zookeeper/bin/zkServer.sh start-foreground")
	return []*types.ServiceDaemon{&types.ServiceDaemon{Command: cmd, Args: args}}
}

func (cb *ZkBoot) WaitForPorts() []int {
	return []int{2181, 2888, 3888}
}

func (cb *ZkBoot) PostBootScripts(currentBoot *types.CurrentBoot) []*types.Script {
	return []*types.Script{}
}

func (cb *ZkBoot) PostBoot(currentBoot *types.CurrentBoot) {
	log.Info("zookeeper: running...")
}

func (cb *ZkBoot) ScheduleTasks(currentBoot *types.CurrentBoot) []*types.Cron {
	return []*types.Cron{}
}

func (cb *ZkBoot) UseConfd() bool {
	return true
}

func (cb *ZkBoot) PreShutdownScripts(currentBoot *types.CurrentBoot) []*types.Script {
	return []*types.Script{}
}

// getMachines return the list of machines that can run zookeeper or an empty list
func getMachines() ([]string, error) {
	metadata, err := fleet.ParseMetadata("zookeeper=true")
	if err != nil {
		panic(err)
	}

	return fleet.GetMachines(fleetEndpoint, metadata)
}

// getNextNodeId returns the next id to use as zookeeper node index
func getNextNodeId(etcdClient *goetcd.Client, nodes []string) int {
	result := 0
	for _, node := range nodes {
		id := etcd.Get(etcdClient, etcdPath+"/"+node+"/id")
		numericId, err := strconv.Atoi(id)
		if id != "" && err == nil && numericId > result {
			result = numericId
		}
	}

	return result + 1
}

// difference get the elements present in the first slice and not in
// the second one returning those elemenets in a new string slice.
func difference(slice1 []string, slice2 []string) []string {
	diffStr := []string{}
	m := map[string]int{}

	for _, s1Val := range slice1 {
		m[s1Val] = 1
	}
	for _, s2Val := range slice2 {
		m[s2Val] = m[s2Val] + 1
	}

	for mKey, mVal := range m {
		if mVal == 1 {
			diffStr = append(diffStr, mKey)
		}
	}

	return diffStr
}
