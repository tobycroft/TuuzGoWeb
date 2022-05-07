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
		//	//c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token")
		//	//c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, PATCH, DELETE")
		//	//c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		//	//c.Header("Access-Control-Allow-Credentials", "true")
		c.Next()
	}
}

func CorsController() gin.HandlerFunc {
	return cors.Default()
}
