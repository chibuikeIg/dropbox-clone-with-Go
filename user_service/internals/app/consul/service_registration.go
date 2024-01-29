package consul_api

import (
	"fmt"
	"log"
	"strconv"
	"user-service/internals/app"

	consulapi "github.com/hashicorp/consul/api"
)

const (
	SERVICE_NAME string = "user-service"
	SERVICE_ID   string = "user-service"
	PING_PATH    string = "api/ping"
	INTERVALS    string = "10s"
	TIMEOUT      string = "30s"
)

func ServiceRegistration() {

	config := consulapi.DefaultConfig()

	consul, err := consulapi.NewClient(config)

	if err != nil {
		log.Printf("Failed to register service: %s", err)
		return
	}

	address := app.HostName()

	port, _ := strconv.Atoi(app.Port()[1:len(app.Port())])

	checkUrl := fmt.Sprintf("http://%s:%v/%s", address, port, PING_PATH)

	registration := &consulapi.AgentServiceRegistration{
		ID:      SERVICE_ID,
		Name:    SERVICE_NAME,
		Port:    port,
		Address: address,
		Check: &consulapi.AgentServiceCheck{
			HTTP:     checkUrl,
			Interval: INTERVALS,
			Timeout:  TIMEOUT,
		},
	}

	regiErr := consul.Agent().ServiceRegister(registration)

	if regiErr != nil {
		log.Printf("Failed to register service: %s ", regiErr)
		return
	} else {
		log.Printf("successfully register service: %s:%v", address, port)
		return
	}
}
