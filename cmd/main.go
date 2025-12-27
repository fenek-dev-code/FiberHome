package main

import (
	"go-fiber/home/config"

	"github.com/gofiber/fiber/v2"
)

func main() {
	config.LoadEnv()
	_ = config.LoadConfigEnv()
	// Application entry point
	app := fiber.New()
}
