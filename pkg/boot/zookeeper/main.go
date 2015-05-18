package main

import (
	"github.com/aledbf/coreos-mesos-zookeeper/pkg/boot"
	"github.com/aledbf/coreos-mesos-zookeeper/pkg/fleet"
	logger "github.com/aledbf/coreos-mesos-zookeeper/pkg/log"
	"github.com/aledbf/coreos-mesos-zookeeper/pkg/os"
	"github.com/aledbf/coreos-mesos-zookeeper/pkg/types"
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
	// lock etcd
	// check if this node is configured
	// if not
	// get fleet nodes with the required role
	// preassing the ids for every node
	// unlock
	// continue
	// go func(){

	// }
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

func getMachines() ([]string, error) {
	metadata, err := fleet.ParseMetadata("role=zookeeper")
	if err != nil {
		panic(err)
	}

	return fleet.GetMachines(fleetEndpoint, metadata)
}
