package controllers

import (
	"log/slog"

	"github.com/gofiber/fiber/v2"
)

type ManageBuilder struct {
	GroupName string
}

func (b *ManageBuilder) Build(app *fiber.App) {
	group := app.Group(b.GroupName)

	group.Get("/", getResults)

	slog.Info("ManageBuilder built", "group", b.GroupName)

}

// getResults sample comment for testing with
// IDE themes
func getResults(c *fiber.Ctx) error {
	slog.Debug("getResults")
	return c.SendString("We are managers...")
}
