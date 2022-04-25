package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/tobycroft/gorose-pro"
)

func IndexController(route fiber.Router) {

	route.All("", index)
	route.All("login", loginss)
	route.All("upload", upload)
	//route.All("register")
}

func index(c *fiber.Ctx) error {
	c.SendString("index")
	return nil
}

func loginss(c *fiber.Ctx) error {
	password := c.Query("password")
	username := c.Query("username")
	json := map[string]string{}
	json["username"] = username
	json["password"] = password
	gorose.Open()
	c.JSON(json)
	return nil
}

func upload(c *fiber.Ctx) error {
	//file, ok := Input.Upload(c)
	//if !ok {
	//	return
	//}
	//fmt.Println(file)
	return nil

}
