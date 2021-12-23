package server

import (
	"strconv"

	"github.com/Pauloo27/k7z/internal/config"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/monitor"
)

func StartHTTPServer() error {
	app := fiber.New()
	app.Get("/admin/dashboard", monitor.New())
	route(app)
	app.Listen(":" + strconv.Itoa(config.Port))
	return nil
}
