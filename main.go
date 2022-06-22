package main

import "fmt"

func main() {

	arr := map[string]interface{}{
		"aaa": "bbb",
		"ccc": "ddd",
		"eee": nil,
	}
	_, ok := arr["fff"]
	fmt.Println(ok)

	//Calc.RefreshBaseNum()
	//mainroute := gin.Default()
	////gin.SetMode(gin.ReleaseMode)
	////gin.DefaultWriter = ioutil.Discard
	//mainroute.SetTrustedProxies([]string{"0.0.0.0/0"})
	//mainroute.SecureJsonPrefix(app_conf.SecureJsonPrefix)
	//route.OnRoute(mainroute)
	//mainroute.Run(":80")

}
