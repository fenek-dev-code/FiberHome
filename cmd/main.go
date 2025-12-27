package main

import (
	"go-fiber/home/config"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	config.LoadEnv()
	_ = config.LoadConfigEnv()
	app := fiber.New()

	app.Use(recover.New())

	app.Listen(":3000")
}
