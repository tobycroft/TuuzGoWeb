package main

import (
	"github.com/gin-gonic/gin"
	"main.go/config/app_conf"
	"main.go/route"
	"main.go/tuuz/Calc"
	"time"
)

func init() {
	var cstZone = time.FixedZone("CST", 8*3600) // 东八
	time.Local = cstZone
}

func main() {
	Calc.RefreshBaseNum()
	mainroute := gin.Default()
	//gin.SetMode(gin.ReleaseMode)
	//gin.DefaultWriter = ioutil.Discard
	mainroute.SetTrustedProxies(nil)
	mainroute.SecureJsonPrefix(app_conf.SecureJsonPrefix)
	route.OnRoute(mainroute)
	mainroute.Run(":80")

}
