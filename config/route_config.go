package config

import (
	"../middleware"
	"github.com/gin-gonic/gin"
)

// SetRoutes .
func SetRoutes(router *gin.Engine) {
	authMiddleware := middleware.Auth()

	router.POST("/login", authMiddleware.LoginHandler)
	auth := router.Group("/auth")
	auth.Use(authMiddleware.MiddlewareFunc())
	{
		auth.GET("/refresh_token", authMiddleware.RefreshHandler)
	}
	SetAPIPaths(router, authMiddleware.MiddlewareFunc())
}
