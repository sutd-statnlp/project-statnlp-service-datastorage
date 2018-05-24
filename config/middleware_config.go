package config

import (
	"../middleware"
	"github.com/gin-gonic/gin"
)

// SetMiddleWares setups middlewares.
func SetMiddleWares(router *gin.Engine) gin.IRoutes {
	iroutes := router.Use(middleware.Static())
	return iroutes
}
