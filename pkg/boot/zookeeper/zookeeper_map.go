package zookeeper

import (
	"strconv"
	"time"

	"github.com/aledbf/coreos-mesos-zookeeper/pkg/etcd"
	"github.com/aledbf/coreos-mesos-zookeeper/pkg/fleet"
	logger "github.com/aledbf/coreos-mesos-zookeeper/pkg/log"
)

const (
	fleetEndpoint = "unix:///var/run/fleet.sock"
)

var (
	log = logger.New()
)

func CheckZkMappingInFleet(etcdPath string, etcdClient *etcd.Client) {
	// check if the nodes with the required role already have the an id. If not
	// get fleet nodes with the required role and preassing the ids for every
	// node in the cluster
	err := etcd.WaitForLock(etcdClient, "/zookeeper/masterLock", 60, 61*time.Second)
	if err != nil {
		panic(err)
	}

	zkNodes := etcd.GetList(etcdClient, etcdPath)
	log.Debugf("zookeeper nodes %v", zkNodes)

	machines, err := getMachines()
	if err != nil {
		panic(err)
	}
	log.Debugf("machines %v", machines)

	if len(zkNodes) == 0 {
		log.Debug("initializing zookeeper cluster")
		for index, newZkNode := range machines {
			etcd.Set(etcdClient, etcdPath+"/"+newZkNode+"/id", strconv.Itoa(index+1), 0)
		}
	} else {
		// we check if some machine in the fleet cluster with the
		// required role is not initialized (no zookeeper node id).
		machinesNotInitialized := difference(machines, zkNodes)
		if len(machinesNotInitialized) > 0 {
			nextNodeID := getNextNodeID(etcdPath, etcdClient, zkNodes)
			for _, zkNode := range machinesNotInitialized {
				etcd.Set(etcdClient, etcdPath+"/"+zkNode+"/id", strconv.Itoa(nextNodeID), 0)
				nextNodeID++
			}
		}
	}

	// release the etcd lock
	etcd.ReleaseLock(etcdClient, "/zookeeper/masterLock")
}

// getMachines return the list of machines that can run zookeeper or an empty list
func getMachines() ([]string, error) {
	metadata, err := fleet.ParseMetadata("zookeeper=true")
	if err != nil {
		panic(err)
	}

	return fleet.GetMachines(fleetEndpoint, metadata)
}

// getNextNodeID returns the next id to use as zookeeper node index
func getNextNodeID(etcdPath string, etcdClient *etcd.Client, nodes []string) int {
	result := 0
	for _, node := range nodes {
		id := etcd.Get(etcdClient, etcdPath+"/"+node+"/id")
		numericID, err := strconv.Atoi(id)
		if id != "" && err == nil && numericID > result {
			result = numericID
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
