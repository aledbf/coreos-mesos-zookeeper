package main

import (
	"github.com/aledbf/coreos-mesos-zookeeper/pkg/boot"
	logger "github.com/aledbf/coreos-mesos-zookeeper/pkg/log"
	"github.com/aledbf/coreos-mesos-zookeeper/pkg/os"
	"github.com/aledbf/coreos-mesos-zookeeper/pkg/types"
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
	cmd, args := os.BuildCommandFromString("mesos-master")
	return []*types.ServiceDaemon{&types.ServiceDaemon{Command: cmd, Args: args}}
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
