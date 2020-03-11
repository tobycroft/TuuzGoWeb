package controller

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"image/png"
	"main.go/tuuz/Base64"
	"main.go/tuuz/Captcha"
	"main.go/tuuz/RET"
)

var username string
var password string
var verify string
var is_verify bool
var is_captcha bool
var captcha string

func LoginController(route *gin.RouterGroup) {
	route.Use(cors.Default())
	route.Use(initialize(), gin.Recovery())

	route.Any("/get_captcha", get_captcha)
	route.Any("/get_captcha_64", get_captcha_64)
	route.Any("/verify_captcha", verify_captcha)
}

func initialize() gin.HandlerFunc {
	return func(c *gin.Context) {

		var is bool
		username, is = c.GetPostForm("username")
		if is == false {
			username, is = c.GetQuery("username")
			if is == false {
				c.JSON(200, RET.Ret_succ(400, "username"))
				c.Abort()
				return
			}
		}
		if len(username) < 8 {
			c.JSON(200, RET.Ret_succ(400, "username"))
			c.Abort()
			return
		}
		password, is = c.GetPostForm("password")
		verify, is_verify = c.GetPostForm("verify")
		captcha, is_captcha = c.GetPostForm("captcha")
		c.Next()
	}
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
	c.JSON(200, RET.Ret_succ(0, ret))
}

func verify_captcha(c *gin.Context) {
	ident, is := c.GetPostForm("ident")
	if is == false {
		c.JSON(200, RET.Ret_succ(400, "ident"))
		c.Abort()
		return
	}
	if is_verify == false {
		c.JSON(200, RET.Ret_succ(400, "verify"))
		c.Abort()
		return
	}
	bol := Captcha.AutoVerify(ident, verify)
	if bol == true {
		c.JSON(200, RET.Ret_succ(0, "验证成功"))
	} else {
		c.JSON(200, RET.Ret_succ(0, "验证失败"))
	}
}
