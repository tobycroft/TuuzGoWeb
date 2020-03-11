package main

import (
	"github.com/gin-gonic/gin"
	"main.go/extend/BiliBili/BiliGenerateRoute"
	"main.go/extend/C2C"
	"main.go/extend/LkTCP"
	"main.go/extend/TCP"
	"main.go/route"
)

func main() {

	mainroute := gin.Default()
	//gin.SetMode(gin.ReleaseMode)
	//gin.DefaultWriter = ioutil.Discard
	route.OnRoute(mainroute)
	mainroute.Run(":80")

}
