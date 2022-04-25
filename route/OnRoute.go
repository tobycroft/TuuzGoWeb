package route

import (
	"github.com/gofiber/fiber/v2"
	recover2 "github.com/gofiber/fiber/v2/middleware/recover"
	v1 "main.go/route/v1"
)

func MainRoute(app *fiber.App) *fiber.App {
	app.Use(recover2.New())
	app.All("", func(c *fiber.Ctx) error {
		c.SendString(c.Path())
		return nil
	})
	version1 := app.Group("v1")
	version1.All("", func(c *fiber.Ctx) error {
		c.SendString(c.Path())
		return nil
	})

	v1.IndexRouter(version1.Group("index"))

	return app
}
