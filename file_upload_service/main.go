package main

import (
	"file-upload-service/internals/app"
	consul_api "file-upload-service/internals/app/consul"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()

	if err != nil {
		panic("error loading .env file. error:" + err.Error())
	}

	consul_api.ServiceRegistration()

	r := gin.Default()

	LoadRoutes(r)

	r.Run(app.Port())
}
