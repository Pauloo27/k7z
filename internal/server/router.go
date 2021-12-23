package server

import (
	"github.com/Pauloo27/k7z/internal/server/controller"
	"github.com/gofiber/fiber/v2"
)

func route(app *fiber.App) {
	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.SendString("hello =)")
	})

	app.Post("/admin/config/reload", controller.ReloadConfig)

	app.Post("/projects/:id/reload", controller.ReloadProject)

}
