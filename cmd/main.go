package main

import (
	"go-fiber/home/config"
	"go-fiber/home/internal/pages"
	"go-fiber/home/pkg/logger"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	slogfiber "github.com/samber/slog-fiber"
)

func main() {
	config.LoadEnv()
	conf := config.LoadConfigEnv()
	logger := logger.NewLogger(conf.LoggerConfig.Level, conf.LoggerConfig.Format)

	app := fiber.New()
	app.Static("/public", "./public")
	app.Use(recover.New())
	app.Use(slogfiber.New(logger))

	pages.NewHanlerPage(app)

	app.Listen(conf.GetServerAddress())
}
