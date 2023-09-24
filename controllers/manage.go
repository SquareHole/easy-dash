package controllers

import (
	"log/slog"

	"github.com/gofiber/fiber/v2"
)

// ManageBuilder is a struct that holds the configutation for the Controller
type ManageBuilder struct {
	GroupName string
}

// Build takes th GroupName from the SysBuilder struct
// creates a group and adds the endpoints to it
func (b *ManageBuilder) Build(app *fiber.App) {

	if b.GroupName == "" {
		panic("GroupName is empty")
	}

	group := app.Group(b.GroupName)

	group.Get("/", b.getResults)

	slog.Info("ManageBuilder built", "group", b.GroupName)

}

// getResults sample comment for testing with
// IDE themes
func (b *ManageBuilder) getResults(c *fiber.Ctx) error {
	slog.Debug("getResults")
	return c.SendString("We are managers...")
}
