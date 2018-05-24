package main

import (
	"./config"
	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	config.SetMiddleWares(router)
	config.SetAPIPaths(router)
	config.SetDbConnection()

	router.Run(":9000")

}
