package controller

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"image/png"
	"main.go/app/v1/index/model/UserMemberModel"
	"main.go/app/v1/user/model/UserTokenModel"
	"main.go/extend/BiliBili/BiliCore"
	"main.go/extend/BiliBili/BiliLogin"
	"main.go/extend/BiliBili/BiliUrl"
	"main.go/tuuz/Base64"
	"main.go/tuuz/Calc"
	"main.go/tuuz/Captcha"
	"main.go/tuuz/Net"
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

	route.Any("/", login)
	route.Any("/bili_captcha", bili_captcha)
	route.Any("/get_captcha", get_captcha)
	route.Any("/get_captcha_64", get_captcha_64)
	route.Any("/verify_captcha", verify_captcha)
	route.Any("/self_login", self_login)
	route.Any("/2", login2)
	route.Any("/ret", login_ret)
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

func bili_captcha(c *gin.Context) {
	query := make(map[string]interface{})
	header := make(map[string]string)
	//fmt.Println(username)
	code, ret, _ := Net.GetCookieAuto(BiliUrl.Captcha, query, BiliCore.Header(header), username)
	if code == 0 {
		c.Header("Content-type", "image/png")
		c.String(200, ret.(string))
	} else {
		c.String(500, ret.(string))
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

func login(c *gin.Context) {
	if len(password) < 6 {
		c.JSON(200, RET.Ret_succ(400, "password_too_short"))
		c.Abort()
		return
	}
	code, ret := BiliLogin.Api_captcha_login(username, password, captcha)
	c.JSON(200, RET.Ret_succ(code, ret))
}

func self_login(c *gin.Context) {
	if len(password) < 6 {
		c.JSON(200, RET.Ret_succ(400, "password_too_short"))
		c.Abort()
		return
	}
	um := UserMemberModel.Api_find(username)
	if len(um) > 0 {
		if um["password"] == password {
			token := Calc.GenerateToken()
			UserTokenModel.Api_insert(username, token, "APP")
			c.JSON(200, RET.Ret_succ(0, map[string]interface{}{"username": username, "token": token, "message": "登录成功，信息更新成功"}))
		} else {
			c.JSON(200, RET.Ret_succ(-1, map[string]interface{}{"message": "用户名或密码错误"}))
		}
	} else {
		c.JSON(200, RET.Ret_succ(-1, map[string]interface{}{"message": "请先正常登陆BiliHP"}))
	}
}

func login2(c *gin.Context) {
	if len(password) < 6 {
		c.JSON(200, RET.Ret_succ(400, "password_too_short"))
		c.Abort()
		return
	}
	ret := BiliLogin.Api_captcha_login2(username, password, captcha)
	c.JSON(200, RET.Ret_succ(0, ret))
}

func login_ret(c *gin.Context) {
	ret, ok := c.GetPostForm("ret")
	if !ok {
		c.JSON(200, RET.Ret_succ(400, "ret"))
		c.Abort()
		return
	}
	code, data := BiliLogin.Api_captcha_login2_ret(username, password, ret)
	c.JSON(200, RET.Ret_succ(code, data))
}
