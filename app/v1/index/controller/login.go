package controller

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"image/png"
	"main.go/tuuz/Base64"
	"main.go/tuuz/Captcha"
	"main.go/tuuz/Input"
	"main.go/tuuz/RET"
)

func LoginController(route *gin.RouterGroup) {
	route.Use(cors.Default())

	route.Any("/get_captcha", get_captcha)
	route.Any("/get_captcha_64", get_captcha_64)
	route.Any("/verify_captcha", verify_captcha)
}

func get_captcha(c *gin.Context) {
	ident, is := c.GetPostForm("ident")
	if is == false {
		c.String(400, "Ident")
		c.Abort()
		return
	}
	img, string := Captcha.ManualCreate(4, ident)
	fmt.Println(string)
	png.Encode(c.Writer, img)
}

func get_captcha_64(c *gin.Context) {
	img, ident := Captcha.AutoCreate()
	fmt.Println(ident)
	b64 := Base64.EncodePng(img)
	ret := make(map[string]string)
	ret["ident"] = ident
	ret["b64"] = b64
	RET.Fail(c, 400, ret, nil)
}

func verify_captcha(c *gin.Context) {
	ident, is := c.GetPostForm("ident")
	if is == false {
		RET.Fail(c, 400, nil, "ident")
		return
	}
	verify, ok := Input.Post("verify", c, false)
	if !ok {
		return
	}
	bol := Captcha.AutoVerify(ident, verify)
	if bol == true {
		RET.Success(c, 0, nil, "验证成功")
	} else {
		RET.Fail(c, 400, nil, "验证失败")
	}
}
