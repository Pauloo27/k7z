package server

import (
	"github.com/Pauloo27/k7z/internal/config"
	"github.com/gofiber/fiber/v2"
)

func route(app *fiber.App) {
	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.SendString("hello =)")
	})

	// TODO: route for reloading the config

	app.Post("/reload/:id", func(ctx *fiber.Ctx) error {
		projectID := ctx.Params("id")
		project, found := config.Projects[projectID]
		if !found {
			return ctx.SendStatus(404)
		}
		return ctx.SendString("hello " + project.Name)
	})
}
