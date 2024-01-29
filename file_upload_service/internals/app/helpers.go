package app

import (
	"os"
)

func Port() string {

	port := os.Getenv("PORT")

	if len(port) == 0 {

		port = "8081"
	}

	return ":" + port
}

func HostName() string {
	hostname, _ := os.Hostname()

	return hostname
}
