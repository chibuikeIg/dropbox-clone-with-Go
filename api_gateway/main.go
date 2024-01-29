package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()

	if err != nil {
		panic("error loading .env file. error:" + err.Error())
	}

	r := gin.Default()

	LoadRoutes(r)

	r.Run()

}
