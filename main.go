package main

import (
	"./config"
	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	config.SetMiddleWares(router)
	config.SetRoutes(router)

	router.Run(":8230")

}
