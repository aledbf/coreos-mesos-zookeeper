package fleet

import (
	"net"
	"net/http"
	"net/url"

	"github.com/coreos/fleet/client"
)

// GetMachines returns the ip address of the nodes with all the specified roles
func GetMachines(endpoint string, metadata map[string][]string) ([]string, error) {
	fleetURL, err := url.Parse(endpoint)
	if err != nil {
		return nil, err
	}

	dialFunc := net.Dial

	// change configuration for socket communication using http
	if fleetURL.Scheme == "unix" {
		sockPath := fleetURL.Path
		fleetURL.Path = ""
		fleetURL.Scheme = "http"
		fleetURL.Host = "domain-sock"

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

	fleetClient, err := client.NewHTTPClient(httpClient, *fleetURL)
	if err != nil {
		return nil, err
	}

	machines, err := fleetClient.Machines()
	if err != nil {
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
