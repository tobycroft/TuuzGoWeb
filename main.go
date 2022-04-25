package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"main.go/config/app_conf"
	"main.go/route"
	"main.go/tuuz/Calc"
)

func main() {
	Calc.RefreshBaseNum()

	conf := fiber.Config{
		AppName:           app_conf.Project,
		EnablePrintRoutes: false,
	}
	app := fiber.New(conf)
	app.Use(logger.New())
	route.MainRoute(app)
	app.Listen(":80")
}
