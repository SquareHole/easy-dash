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

// RS is a global variable that holds the scheduler instance
var RS *scheduling.Scheduler

// Serve starts the web server and listens for incoming requests
func Serve(s *scheduling.Scheduler) error {

	// Set the global scheduler instance
	RS = s

	logger := slog.Default()

	// Create a new Fiber app
	app := fiber.New()

	// Add middleware
	app.Use(slogfiber.New(logger))
	app.Use(cors.New(cors.Config{
		AllowMethods:     "GET,POST,PUT,DELETE",
		AllowCredentials: true,
		AllowOrigins:     "*",
	}))

	// Add a hook to log incoming requests
	app.Hooks().OnRoute(func(r fiber.Route) error {
		slog.Debug("Route", "method", r.Method, "path", r.Path)
		return nil
	})

	// Bind the routes to the app
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

	// Start listening for incoming requests
	return app.Listen(port)
}

// bindRoutes binds the route handlers to the app
func bindRoutes(app *fiber.App) {
	var wg sync.WaitGroup

	// Create an array of controller builders
	var builders = []controllers.Builder{
		&controllers.ConfigBuilder{GroupName: "/_config", Scheduler: RS},
		&controllers.SysBuilder{GroupName: "/_sys"},
		&controllers.ManageBuilder{GroupName: "/manage"},
	}

	for _, builder := range builders {
		// Increment the WaitGroup counter.
		wg.Add(1)

		// Launch a goroutine to register the route handlers.
		go func(builder controllers.Builder) {
			defer wg.Done()
			builder.Build(app)
		}(builder)
	}

	// Wait for all the goroutines to finish
	wg.Wait()
}
