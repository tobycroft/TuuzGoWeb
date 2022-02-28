package main

import (
	"github.com/gin-gonic/gin"
	"main.go/config/app_conf"
	"main.go/route"
)

func main() {

	mainroute := gin.Default()
	//gin.SetMode(gin.ReleaseMode)
	//gin.DefaultWriter = ioutil.Discard
	mainroute.Static("/", app_conf.FileSavePath)
	mainroute.SetTrustedProxies(nil)
	mainroute.SecureJsonPrefix(app_conf.SecureJsonPrefix)
	route.OnRoute(mainroute)
	mainroute.Run(":80")

}
