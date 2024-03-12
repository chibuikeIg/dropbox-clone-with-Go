package consul_api

import (
	"errors"

	consulapi "github.com/hashicorp/consul/api"
)

func ServiceDiscovery(serviceName string) (string, int, error) {

	config := consulapi.DefaultConfig()

	consul, err := consulapi.NewClient(config)

	if err != nil {
		return "", 0, errors.New("internal server error occured")
	}

	services, err := consul.Agent().Services()

	if err != nil {
		return "", 0, errors.New("internal server error occured")
	}

	service := services[serviceName+"-service"]

	if service == nil {
		return "", 0, errors.New("service not found")
	}

	address := service.Address
	port := service.Port

	return address, port, nil
}
