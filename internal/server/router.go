package server

import (
	"github.com/Pauloo27/k7z/internal/config"
	"github.com/Pauloo27/logger"
	"github.com/gofiber/fiber/v2"
)

func route(app *fiber.App) {
	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.SendString("hello =)")
	})

	app.Post("/admin/config/reload", func(ctx *fiber.Ctx) error {
		secret := ctx.Get("Secret")
		if secret == "" {
			return ctx.SendStatus(fiber.StatusUnauthorized)
		}
		if secret != config.AdminSecret {
			return ctx.SendStatus(fiber.StatusForbidden)
		}
		err := config.LoadConfig()
		if err != nil {
			return ctx.SendStatus(fiber.StatusInternalServerError)
		}
		return ctx.SendStatus(fiber.StatusOK)
	})

	app.Post("/projects/:id/reload", func(ctx *fiber.Ctx) error {
		projectID := ctx.Params("id")
		secret := ctx.Get("Secret")
		if secret == "" {
			return ctx.SendStatus(fiber.StatusUnauthorized)
		}

		project, found := config.Projects[projectID]
		if !found {
			return ctx.SendStatus(fiber.StatusNotFound)
		}
		if secret != project.Secret {
			return ctx.SendStatus(fiber.StatusForbidden)
		}
		go func() {
			err := project.Reload()
			if err != nil {
				logger.Errorf("project %s (%s) failed to reload: %v", project.Name, project.ID, err)
				return
			}
			logger.Successf("project %s (%s) successfully reloaded", project.Name, project.ID)
		}()
		return ctx.SendStatus(fiber.StatusOK)
	})
}
