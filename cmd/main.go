package main

import (
	"go-fiber/home/config"
	"go-fiber/home/internal/pages"
	"go-fiber/home/pkg/logger"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/template/html/v2"
	slogfiber "github.com/samber/slog-fiber"
)

func main() {
	config.LoadEnv()
	cfg := config.LoadConfigEnv()
	logger := logger.NewLogger(cfg.LoggerConfig.Level, cfg.LoggerConfig.Format)
	tmplEngine := html.New("./views/html", ".html")

	app := fiber.New(fiber.Config{
		Views: tmplEngine,
	})

	// Serve static files (CSS, JS, images) from the views folder at /static
	// This makes files in ./views/css available at /static/css/...
	app.Static("/static", "./views")
	pages.NewHanlerPage(app)
	app.Use(recover.New())
	app.Use(slogfiber.New(logger))

	app.Listen(":" + strconv.Itoa(cfg.Port))
}
