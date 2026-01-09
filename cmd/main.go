package main

import (
	"go-fiber/home/config"
	"go-fiber/home/internal/pages"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	config.LoadEnv()
	cfg := config.LoadConfigEnv()
	app := fiber.New()
	pages.NewHanlerPage(app)
	app.Use(recover.New())

	app.Listen(":" + strconv.Itoa(cfg.Port))
}
