package consul_api

import (
	"fmt"
	"net/http"

	consulapi "github.com/hashicorp/consul/api"
)

func ServiceDiscovery(serviceName string, reqPath string) (string, int) {

	config := consulapi.DefaultConfig()

	consul, err := consulapi.NewClient(config)

	if err != nil {
		return "", http.StatusInternalServerError
	}

	services, err := consul.Agent().Services()

	if err != nil {
		return "", http.StatusInternalServerError
	}

	service := services[serviceName+"-service"]

	if service == nil {
		return "", http.StatusNotFound
	}

	address := service.Address
	port := service.Port

	// clear request path
	rp := reqPath[1:]

	return fmt.Sprintf("http://%s:%v/api/v1/%s", address, port, rp), 0
}
