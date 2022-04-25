package v1

import (
	"github.com/gofiber/fiber/v2"
	"main.go/app/v1/index/controller"
)

func IndexRouter(route fiber.Router) {
	route.All("/", func(context *fiber.Ctx) error {
		context.SendString(context.Path())
		return nil
	})

	controller.IndexController(route.Group("index"))
	controller.LoginController(route.Group("login"))
}
