package main

import (
	"file-upload-service/internals/app"
	consul_api "file-upload-service/internals/app/consul"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()

	if err != nil {
		panic("error loading .env file. error:" + err.Error())
	}

	// consul service registration
	consul_api.ServiceRegistration()

	// REST endpoints registration
	r := gin.Default()
	LoadRoutes(r)

	if len(os.Args) >= 2 {
		os.Exit(0)
	}

	r.Run(app.Port())
}
