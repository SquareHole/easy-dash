package controllers

import (
	"log/slog"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/squarehole/easydash/internal/data"
	"github.com/squarehole/easydash/internal/scheduling"
)

// ConfigBuilder is a struct that holds the configutation for the Controller
type ConfigBuilder struct {
	GroupName string
	Scheduler *scheduling.Scheduler
}

// Build takes th GroupName from the SysBuilder struct
// creates a group and adds the endpoints to it
func (b *ConfigBuilder) Build(app *fiber.App) {

	if b.GroupName == "" {
		panic("GroupName is empty")
	}

	group := app.Group(b.GroupName)
	group.Get("/", b.getConfig)
	group.Get("/schedules", b.getScheduledJobs)

	group.Post("/schedule/:duration", b.addSchedule)
	group.Put("/schedule/:jobId", b.updateSchedule)

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
	slog.Debug("getScheduledJobs")
	jobs := &b.Scheduler.Schedules
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"jobs": jobs})
}

func (b *ConfigBuilder) stopScheduledJob(c *fiber.Ctx) error {
	slog.Debug("stopScheduledJob")
	jobId, err := strconv.Atoi(c.Params("jobId"))
	if err != nil {
		slog.Error("Error while converting jobId", "error", err.Error())
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	slog.Info("Stopping job", "jobId", jobId)
	b.Scheduler.StopJobById(jobId)
	return c.SendStatus(fiber.StatusOK)
}

func (b *ConfigBuilder) addSchedule(c *fiber.Ctx) error {

	duration := c.Params("duration")

	b.Scheduler.AddSchedule("Test", "@every "+duration, func() {
		slog.Debug("Running job added later", "duration", duration, "time", time.Now())
	})
	return c.SendStatus(fiber.StatusOK)
}

func (b *ConfigBuilder) updateSchedule(c *fiber.Ctx) error {
	jobId, err := strconv.Atoi(c.Params("jobId"))
	if err != nil {
		slog.Error("JobId not supplied")
		return c.SendStatus(fiber.StatusBadRequest)
	}

	slog.Debug("Updating job", "jobId", jobId)
	return c.SendStatus(fiber.StatusOK)
}
