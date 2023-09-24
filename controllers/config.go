package controllers

import (
	"log/slog"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/squarehole/easydash/data"
	"github.com/squarehole/easydash/scheduling"
)

// ConfigBuilder is a struct that holds the configutation for the Controller
type ConfigBuilder struct {
	GroupName string
	Scheduler *scheduling.Scheduler
}

// Build takes th GroupName from the SysBuilder struct
// creates a group and adds the endpoints to it
func (b *ConfigBuilder) Build(app *fiber.App) {

	group := app.Group(b.GroupName)
	group.Get("/", b.getConfig)
	group.Get("/schedules", b.getScheduledJobs)

	group.Delete("/schedule/:jobId", b.stopScheduledJob)

	slog.Info("ConfigBuilder built", "group", b.GroupName)
}

func (b *ConfigBuilder) getConfig(c *fiber.Ctx) error {

	data, err := data.GetAllConfigs()
	if err != nil {
		slog.Error("Error while getting configs", "error", err.Error())
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.Status(fiber.StatusOK).JSON(data)
}

func (b *ConfigBuilder) getScheduledJobs(c *fiber.Ctx) error {
	slog.Info("getScheduledJobs")
	jobs := &b.Scheduler.Schedules
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"jobs": jobs})
}

func (b *ConfigBuilder) stopScheduledJob(c *fiber.Ctx) error {
	slog.Info("stopScheduledJob")
	jobId, err := strconv.Atoi(c.Params("jobId"))
	if err != nil {
		slog.Error("Error while converting jobId", "error", err.Error())
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	slog.Info("Stopping job", "jobId", jobId)
	b.Scheduler.StopJobById(jobId)
	return c.SendStatus(fiber.StatusOK)
}
