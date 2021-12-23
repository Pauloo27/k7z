package server

import (
	"strconv"

	"github.com/Pauloo27/k7z/internal/config"
	"github.com/gofiber/fiber/v2"
)

func StartHTTPServer() error {
	app := fiber.New()
	route(app)
	app.Listen(":" + strconv.Itoa(config.Port))
	return nil
}
