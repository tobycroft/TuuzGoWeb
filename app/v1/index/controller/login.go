package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func LoginController(route fiber.Router) {
	route.Use(cors.New())

	route.All("get_captcha", get_captcha)
	route.All("get_captcha_64", get_captcha_64)
	route.All("verify_captcha", verify_captcha)
}

func get_captcha(c *fiber.Ctx) error {
	//ident, is := c.GetPostForm("ident")
	//if is == false {
	//	c.String(400, "Ident")
	//	c.Abort()
	//	return
	//}
	//img, string := Captcha.ManualCreate(4, ident)
	//fmt.Println(string)
	//png.Encode(c.Writer, img)
	return nil

}

func get_captcha_64(c *fiber.Ctx) error {
	//img, ident := Captcha.AutoCreate()
	//fmt.Println(ident)
	//b64 := Base64.EncodePng(img)
	//ret := make(map[string]string)
	//ret["ident"] = ident
	//ret["b64"] = b64
	//RET.Fail(c, 400, ret, nil)
	return nil

}

func verify_captcha(c *fiber.Ctx) error {
	//ident, is := c.GetPostForm("ident")
	//if is == false {
	//	RET.Fail(c, 400, nil, "ident")
	//	return
	//}
	//verify, ok := Input.Post("verify", c, false)
	//if !ok {
	//	return
	//}
	//bol := Captcha.AutoVerify(ident, verify)
	//if bol == true {
	//	RET.Success(c, 0, nil, "验证成功")
	//} else {
	//	RET.Fail(c, 400, nil, "验证失败")
	//}
	return nil

}
