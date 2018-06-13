package middleware

import (
	"time"

	jwt "github.com/appleboy/gin-jwt"
	"github.com/gin-gonic/gin"
)

// User .
type User struct {
	UserName  string
	FirstName string
	LastName  string
}

// Auth .
func Auth() *jwt.GinJWTMiddleware {
	authMiddleware := &jwt.GinJWTMiddleware{
		Realm:         "test zone",
		Key:           []byte("SECRET_KEY"),
		Timeout:       time.Hour,
		MaxRefresh:    time.Hour,
		Authenticator: handleAuthenticator,
		Authorizator:  handleAuthorizator,
		Unauthorized:  handleUnauthorized,
		TokenLookup:   "header:Authorization",
		TokenHeadName: "Bearer",
		TimeFunc:      time.Now,
	}
	return authMiddleware
}

func handleAuthenticator(userID string, password string, c *gin.Context) (interface{}, bool) {
	if userID == "admin" && password == "admin" {
		return &User{
			UserName:  userID,
			LastName:  "Admin",
			FirstName: "Admin",
		}, true
	}
	return nil, false
}

func handleAuthorizator(user interface{}, c *gin.Context) bool {
	if v, ok := user.(string); ok && v == "admin" {
		return true
	}
	return false
}

func handleUnauthorized(c *gin.Context, code int, message string) {
	c.JSON(code, gin.H{
		"code":    code,
		"message": "Unauthorized",
	})
}
