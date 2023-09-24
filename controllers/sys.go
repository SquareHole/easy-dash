package controllers

import (
	"time"

	"log/slog"

	"github.com/gofiber/fiber/v2"
)

type SysBuilder struct {
	GroupName string
}

func (b *SysBuilder) Build(app *fiber.App) {

	group := app.Group(b.GroupName)
	group.Get("/poke", poke)

	slog.Info("SysBuilder built", "group", b.GroupName)
}

func poke(c *fiber.Ctx) error {
	slog.Info("Poked...", "when", time.Now())
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Poked"})
}
