package fleet

import (
	"net"
	"net/http"
	"net/url"

	"github.com/coreos/fleet/client"
)

// GetMachines returns the ip address of the nodes with all the specified roles
func GetMachines(endpoint string, metadata map[string][]string) ([]string, error) {
	fleetUrl, err := url.Parse(endpoint)
	if err != nil {
		return nil, err
	}

	dialFunc := net.Dial

	// change configuration for socket communication using http
	if fleetUrl.Scheme == "unix" {
		sockPath := fleetUrl.Path
		fleetUrl.Path = ""
		fleetUrl.Scheme = "http"
		fleetUrl.Host = "domain-sock"

		dialFunc = func(network, addr string) (net.Conn, error) {
			return net.Dial("unix", sockPath)
		}
	}

	httpClient := &http.Client{
		Transport: &http.Transport{
			Dial:              dialFunc,
			DisableKeepAlives: true,
		},
	}

	fleetClient, err := client.NewHTTPClient(httpClient, *fleetUrl)
	if err != nil {
		return nil, err
	}

	machines, err := fleetClient.Machines()
	if err != nil {
		return nil, err
	}

	machineList := make([]string, 0)
	for _, m := range machines {
		if hasMetadata(m, metadata) {
			machineList = append(machineList, m.PublicIP)
		}
	}

	return machineList, nil
}
