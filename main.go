package main

import (
	"github.com/gin-gonic/gin"
	"github.com/tobycroft/Calc"
	"main.go/config/app_conf"
	"main.go/extend/PortFinder"
	"main.go/route"
	"os"
)

func init() {
	if app_conf.TestMode == false {
		s, err := os.Stat("./log/")

		if err != nil {
			os.Mkdir("./log", 0755)
		} else if s.IsDir() {
			os.Mkdir("./log", 0755)
		}
	}
}

func main() {

	Calc.RefreshBaseNum()
	mainroute := gin.Default()
	//gin.SetMode(gin.ReleaseMode)
	//gin.DefaultWriter = ioutil.Discard
	mainroute.SetTrustedProxies([]string{"0.0.0.0/0"})
	mainroute.SecureJsonPrefix(app_conf.SecureJsonPrefix)
	route.OnRoute(mainroute)
	if !PortFinder.PortInUse(80) {
		mainroute.Run(":80")
	} else {
		mainroute.Run(":81")
	}

}
