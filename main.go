package main

import (
	"crypto"
	"fmt"
	"main.go/tuuz/RSA"
)

//func init() {
//	time.Local = app_conf.TimeZone
//	if app_conf.TestMode == false {
//		s, err := os.Stat("./log/")
//
//		if err != nil {
//			os.Mkdir("./log", 0755)
//		} else if s.IsDir() {
//			os.Mkdir("./log", 0755)
//		}
//	}
//}
//
//func main2() {
//
//	Calc.RefreshBaseNum()
//	mainroute := gin.Default()
//	//gin.SetMode(gin.ReleaseMode)
//	//gin.DefaultWriter = ioutil.Discard
//	mainroute.SetTrustedProxies([]string{"0.0.0.0/0"})
//	mainroute.SecureJsonPrefix(app_conf.SecureJsonPrefix)
//	route.OnRoute(mainroute)
//	mainroute.Run(":80")
//	mainroute.Run(":81")
//
//}

func main() {
	fmt.Println(RSA.RsaSign("test", "-----BEGIN RSA PRIVATE KEY-----\n-----BEGIN RSA PRIVATE KEY-----\nMIIEowIBAAKCAQEAkYQwdEo3UAyBlUGdmfDbt1IiCm5/e1cVzkstle7mLmkGWSnG\nLQbBxlAbKqh13jTcAr5PGlCt1vXyBBCi6UrOwBYcUeKehXNZ85/FCR6LG96cUCe6\nLWkoo1Wr1xu1Wl8fTDqlqAkQar6CbZFAn1D7vNCsIqbPsOspdOVWL4YCnCIBLZFn\n6HoL9XPBBIRhFH6Bg4FEam6A0Qx4EfoJczNpMKzTeGHphF0iqhn29bVd8nOUgVvb\nmBBzftuVHeYwJJ62rRH05lvn4MWS/wAU+l6QZSONXgCTb6oW1LSSotnr1s6zTXUq\nWAefHMz4S0labVS1/KzLiXslKw/jxoIMqGd9awIDAQABAoIBAAegGSVaVSBEvlx+\nOXd0P7Pab7CDHB+nFAoFtoFBcnnnjeEJsIMpPTLaohHP6Pi9HGE6tYOcprZL4a0X\nJGAObKTh8XMdoHNAGVoMDUqlC6v7WKoxfQpmcFK6RNptbwAFWcaDtWO+jLWv2ppu\nsQMV9ZslGM3YDaxyAbj/mU7h8Sx/lSDT9E6nsKTr1YnF6Pt2owUWhVx8TmYWQ74Z\nGRoE4kBWdN2jUBsYwRgygDlIz3eRb+IrzVCZ1CrFwuselBSbudza26VyOVtLI27W\nbIqsJbvslIrG45llsEBbBKNAvFMpMrZ5iwy4iwnBXpFe0uIwihNX3UcQOSQtoRwZ\n/LbwYUECgYEAzQxAh4+eZbyaDXal1CY6Kono+hdks2OrOz/9kPoqumU2aQh9YzGj\n2QrfUGatVJJeAE1CHtb7JVIWE2c2aL02pEE9fD6i+f2x8FTWo600x5REZ4tUq4PK\nrGa3FJ3j29w42E//wn32QJuYzdzIoRqFGGsUqAD05+8t+59wUQeJbiUCgYEAtazy\n9BZdyeoOKpMQhde4OlQfpa5TBbsRRuNuLO5oSjh0aMTQrERDJ9C2fMnST1/YtT6P\nVvaP4Y2+8t4ExjqyC/0R8S2hrRIxdk2XQgYP8O+J1iG6mD8COKdAKme5Lpzjn40H\nkMmy6PCPox0rTPhAxb7Ff+Ad8yFPJ2fssvR8gE8CgYAaDMJt9nsDJdDUgQxURX8C\nRH8KtsoeWD06y8hlEMM45v7gnSmA3p+YRV9VkZXXtrimhSovTNSfSyxuzqAE8nTs\nrTUogbSm2eovRDf4l1qrOFTwaq0ZFSc8e9gqkFyQZAv8vz0Y8nPEhYAGN1Rt0zax\ngXkgVu7GQIaw/vJ/+Nsm2QKBgBWLeRRxQpYbZ6qs3hzBRFvGdipTzgyz7oyVlA9I\nBp4mq8dw774+Kiiim8GYvZQkLbLwxFbvzohVIvvyUGaht1Oso2ASpUW9DpiVAwcV\naPaNsa7vOQ2tCzpkuztMKa3ZdTUKqIcHJGxzetVBNE1gc23//bP4hBS9MoHd1Tgz\nvkAJAoGBALSk5XU5YCDgLtiATYZ9ABeHOnT7DGl6bof3gl75hkRxJacEWeBTJccR\nhyK3ITV2mIP4mygcS6DDViW/bc/XJb6Gf12SVYlsA4gv5N2Wsk7N1oUcfT+YSrJC\nBz9nL+D3IDx5NGoowAzpTOfRUZEAlK3ZbmBHqVvXi+GKNiLummdB\n-----END RSA PRIVATE KEY-----",
		crypto.SHA256))
}
