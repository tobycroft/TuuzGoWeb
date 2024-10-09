package route

import (
	"github.com/gin-gonic/gin"
	Net "github.com/tobycroft/TuuzNet"
	v1 "main.go/route/v1"
)

func OnRoute(router *gin.Engine) {
	router.Any("/", func(c *gin.Context) {
		ws := Net.WsServer{}
		ws.NewServer(c.Writer, c.Request, c.Writer.Header())
	})
	version1 := router.Group("/v1")
	{
		version1.Use(func(context *gin.Context) {
		}, gin.Recovery())
		version1.Any("/", func(context *gin.Context) {
			context.String(0, version1.BasePath())
		})
		index := version1.Group("index")
		{
			v1.IndexRouter(index)
		}

	}
}
