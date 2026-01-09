package main

import (
	"go-fiber/home/config"
	"go-fiber/home/internal/pages"
	"go-fiber/home/pkg/logger"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	slogfiber "github.com/samber/slog-fiber"
)

func main() {
	config.LoadEnv()
	cfg := config.LoadConfigEnv()
	logger := logger.NewLogger(cfg.LoggerConfig.Level, cfg.LoggerConfig.Format)

	app := fiber.New()
	app.Static("/public", "./public")
	app.Use(recover.New())
	app.Use(slogfiber.New(logger))

	pages.NewHanlerPage(app)

	app.Listen(":" + strconv.Itoa(cfg.Port))
}
