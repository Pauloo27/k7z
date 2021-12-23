package controller

import (
	"github.com/Pauloo27/k7z/internal/config"
	"github.com/Pauloo27/logger"
	"github.com/gofiber/fiber/v2"
)

func ReloadConfig(ctx *fiber.Ctx) error {
	secret := ctx.Get("Authorization")
	if secret == "" {
		logger.Info("Got a request missing authorization header")
		return ctx.SendStatus(fiber.StatusUnauthorized)
	}
	if secret != config.AdminSecret {
		logger.Info("Got a request with an invalid secret")
		return ctx.SendStatus(fiber.StatusForbidden)
	}
	err := config.LoadConfig()
	if err != nil {
		logger.Error("Got error when reloading config:", err)
		return ctx.SendStatus(fiber.StatusInternalServerError)
	}
	logger.Success("Config reloaded")
	return ctx.SendStatus(fiber.StatusOK)
}
