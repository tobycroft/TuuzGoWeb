package main

import (
	"github.com/gofiber/fiber/v2"
	"main.go/route"
	"main.go/tuuz/Calc"
)

func main() {
	Calc.RefreshBaseNum()

	app := fiber.New()
	route.MainRoute(app)
	app.Listen(":80")
}
