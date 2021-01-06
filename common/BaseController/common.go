package BaseController

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"main.go/config/app_conf"
	"net/http"
)

func CommonController() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("S-P-I", c.ClientIP())
		c.Header("S-P-P", app_conf.Project)
		c.Header("S-P-M", app_conf.AppMode)
		//c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token")
		//c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, PATCH, DELETE")
		//c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		//c.Header("Access-Control-Allow-Credentials", "true")
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		c.Next()
	}
}

func CorsController() gin.HandlerFunc {
	return cors.Default()
}
