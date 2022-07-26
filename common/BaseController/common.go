package BaseController

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CommonController() gin.HandlerFunc {
	return func(c *gin.Context) {
		header_handler(c)
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}
		c.Next()
	}
}

func CorsController() gin.HandlerFunc {
	conf := cors.DefaultConfig()
	conf.AllowAllOrigins = true
	conf.AllowHeaders = []string{"*"}
	conf.AllowCredentials = true
	conf.AllowMethods = []string{"POST", "GET", "OPTIONS", "PUT", "PATCH", "DELETE"}
	conf.ExposeHeaders = []string{"*"}
	return cors.New(conf)
}
