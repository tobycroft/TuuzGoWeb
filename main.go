package main

import "fmt"

func main() {

	arr := map[string]interface{}{
		"aaa": "bbb",
		"ccc": "ddd",
		"eee": nil,
	}
	ass(&arr)
	fmt.Println(arr)

	//Calc.RefreshBaseNum()
	//mainroute := gin.Default()
	////gin.SetMode(gin.ReleaseMode)
	////gin.DefaultWriter = ioutil.Discard
	//mainroute.SetTrustedProxies([]string{"0.0.0.0/0"})
	//mainroute.SecureJsonPrefix(app_conf.SecureJsonPrefix)
	//route.OnRoute(mainroute)
	//mainroute.Run(":80")

}

func ass(arr *map[string]interface{}) {
	temp := *arr
	temp["eee"] = "fff"
	arr = &temp
}
