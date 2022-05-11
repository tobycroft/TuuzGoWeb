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
	return cors.Default()
}
