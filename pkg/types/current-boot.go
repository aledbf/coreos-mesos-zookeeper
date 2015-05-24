package types

import (
	"net"
	"time"

	"github.com/aledbf/coreos-mesos-zookeeper/pkg/etcd"
)

// CurrentBoot information about the boot
// process related to the component
type CurrentBoot struct {
	ConfdNodes []string
	EtcdClient *etcd.Client
	EtcdPath   string
	EtcdPort   int
	Host       net.IP
	Port       int
	Timeout    time.Duration
	TTL        time.Duration
}
