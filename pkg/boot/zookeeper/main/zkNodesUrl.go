package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"text/template"

	"github.com/aledbf/coreos-mesos-zookeeper/pkg/etcd"
	logger "github.com/aledbf/coreos-mesos-zookeeper/pkg/log"
)

const (
	zkURLTemplate = "{{ range $index, $node := .nodes }}{{ if $index }},{{ end }}{{ $node }}:3888{{ end }}"
	etcdPath      = "/zookeeper/nodes"
)

var (
	log = logger.New()
)

func init() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: [options]\n\n")
		flag.PrintDefaults()
	}
}

func main() {

	etcdPeers := flag.String("peers", "localhost:4001", "etcd peer/s. For more than one node use comma as separator: localhost:4001,localhost:4001)")

	flag.Parse()

	if flag.NFlag() < 1 {
		flag.Usage()
		os.Exit(1)
	}

	etcdClient := etcd.NewClient(getHTTPEtcdUrls(*etcdPeers))

	zkNodes := etcd.GetList(etcdClient, etcdPath)

	data := make(map[string]interface{})
	data["nodes"] = zkNodes

	t := template.New("zkTemplate")
	t, err := t.Parse(zkURLTemplate)
	if err != nil {
		panic(err)
	}

	err = t.Execute(os.Stdout, data)
	if err != nil {
		fmt.Fprint(os.Stdout, "")
	}

}

// getEtcdHosts returns an array of urls that contains at least one host
func getHTTPEtcdUrls(etcdPeers string) []string {
	hosts := strings.Split(etcdPeers, ",")
	result := []string{}
	for _, host := range hosts {
		result = append(result, "http://"+host)
	}

	return result
}
