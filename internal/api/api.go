package api

import (
	"os"
	"sync"

	"log/slog"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	slogfiber "github.com/samber/slog-fiber"
	"github.com/squarehole/easydash/internal/controllers"
	"github.com/squarehole/easydash/internal/scheduling"
)

var RS *scheduling.Scheduler

func Serve(s *scheduling.Scheduler) error {

	RS = s
	//w := os.Stderr

	// slog.SetDefault(slog.New(
	// 	tint.NewHandler(w, &tint.Options{
	// 		Level:      slog.LevelDebug,
	// 		TimeFormat: time.Kitchen,
	// 	}),
	// ))

	logger := slog.Default()

	app := fiber.New()

	// Add middleware
	app.Use(slogfiber.New(logger))
	app.Use(cors.New(cors.Config{
		AllowMethods:     "GET,POST,PUT,DELETE",
		AllowCredentials: true,
		AllowOrigins:     "*",
	}))

	app.Hooks().OnRoute(func(r fiber.Route) error {
		slog.Debug("Route", "method", r.Method, "path", r.Path)
		return nil
	})

	bindRoutes(app)

	// Get the listening port from the environment
	port := os.Getenv("PORT")

	// when port is empty, use the default port of :3000
	if port == "" {
		port = ":3000"
	}

	// Ensure that port starts with a ":"
	if port[0] != ':' {
		port = ":" + port
	}

	return app.Listen(port)
}

func bindRoutes(app *fiber.App) {
	var wg sync.WaitGroup

	var builders = []controllers.Builder{
		&controllers.ConfigBuilder{GroupName: "/_config", Scheduler: RS},
		&controllers.SysBuilder{GroupName: "/_sys"},
		&controllers.ManageBuilder{GroupName: "/manage"},
	}

	for _, builder := range builders {
		// Increment the WaitGroup counter.
		wg.Add(1)

		// Launch a goroutine to registrer the route handlers.
		go func(builder controllers.Builder) {
			defer wg.Done()
			builder.Build(app)
		}(builder)
	}

	wg.Wait()
}
