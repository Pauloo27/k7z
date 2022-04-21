package controller

import (
	"github.com/Pauloo27/k7z/internal/config"
	"github.com/Pauloo27/logger"
	"github.com/gofiber/fiber/v2"
)

func ReloadProject(ctx *fiber.Ctx) error {
	projectID := ctx.Params("id")

	project, found := config.Projects[projectID]
	// if not found stuff is done only after the auth is checked to avoid leaking
	// information about the existence of projects.

	secretHeader := project.SecretHeader
	if secretHeader == "" {
		secretHeader = "Authorization"
	}

	secret := ctx.Get(secretHeader)

	if secret == "" {
		logger.Info("Got a request missing authorization header")
		return ctx.SendStatus(fiber.StatusUnauthorized)
	}

	if secret != project.Secret {
		logger.Info("Got a request with an invalid secret")
		return ctx.SendStatus(fiber.StatusForbidden)
	}

	if !found {
		logger.Info("Cannot found project with ID", projectID)
		return ctx.SendStatus(fiber.StatusNotFound)
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
}
