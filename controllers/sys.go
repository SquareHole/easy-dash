package controllers

import (
	"time"

	"log/slog"

	"github.com/gofiber/fiber/v2"
)

// SysBuilder is a struct that holds the configutation for the Controller
type SysBuilder struct {
	GroupName string
}

// Build takes th GroupName from the SysBuilder struct
// creates a group and adds the endpoints to it
func (b *SysBuilder) Build(app *fiber.App) {

	group := app.Group(b.GroupName)
	group.Get("/poke", b.poke)

	slog.Info("SysBuilder built", "group", b.GroupName)
}

// poke is a test endpoint to see if the server is running
// It responds with a 200 and a JSON message containing Ouch...
func (b *SysBuilder) poke(c *fiber.Ctx) error {
	slog.Info("Poked...", "when", time.Now())
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Ouch..."})
}
