package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"main.go/config/app_conf"
	"main.go/route"
	"main.go/tuuz/Calc"
	"main.go/tuuz/Date"
)

func main() {

	data := Date.Offset_week1st(0)
	fmt.Println(Date.Offset_format_second(data))
	return
	Calc.RefreshBaseNum()
	mainroute := gin.Default()
	//gin.SetMode(gin.ReleaseMode)
	//gin.DefaultWriter = ioutil.Discard
	mainroute.SetTrustedProxies([]string{"0.0.0.0/0"})
	mainroute.SecureJsonPrefix(app_conf.SecureJsonPrefix)
	route.OnRoute(mainroute)
	mainroute.Run(":80")

}
