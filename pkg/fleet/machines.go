package fleet

import (
	"net/http"
	"time"

	logger "github.com/aledbf/coreos-mesos-zookeeper/pkg/log"
	"github.com/coreos/fleet/etcd"
	"github.com/coreos/fleet/registry"
)

var log = logger.New()

// GetMachinesWithMetadata returns the ip address of the nodes with all the specified roles
func GetMachinesWithMetadata(url []string, metadata map[string][]string) ([]string, error) {
	etcdClient, err := etcd.NewClient(url, &http.Transport{}, time.Second)
	if err != nil {
		log.Debugf("error creating new fleet etcd client: %v", err)
		return nil, err
	}

	fleetClient := registry.NewEtcdRegistry(etcdClient, "/_coreos.com/fleet/")
	machines, err := fleetClient.Machines()
	if err != nil {
		log.Debugf("error creating new fleet etcd client: %v", err)
		return nil, err
	}

	var machineList []string
	for _, m := range machines {
		if hasMetadata(m, metadata) {
			machineList = append(machineList, m.PublicIP)
		}
	}

	return machineList, nil
}
