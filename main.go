package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"sync"

	"log/slog"

	"github.com/joho/godotenv"
	"github.com/lmittmann/tint"
	"github.com/squarehole/easydash/internal/api"
	"github.com/squarehole/easydash/internal/data"
	"github.com/squarehole/easydash/internal/scheduling"
)

func init() {
	// Read Configuration data from the .env file in the project
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file : " + err.Error())
	}
}

// main is the entry point of the application
func main() {

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()

		setupLogger()

		// Initialize the database
		err := data.InitDatabase()
		if err != nil {
			// Log the error if database initialization fails
			slog.Error("Error while initializing the database", "error", err.Error())
			panic(err)
		}
	}()

	// Create a new scheduler
	rs := scheduling.New()

	// Start the web server, passing in the scheduler
	log.Fatal(api.Serve(rs))
}

func setupLogger() {

	// Check if logging to file is enabled
	logToFile, err := strconv.ParseBool(os.Getenv("LOG_TO_FILE"))
	if err != nil {
		logToFile = false
	}

	if logToFile {

		// If logging to file is enabled, create a new log file
		fileName := os.Getenv("LOG_FILE")
		fmt.Println(fileName)
		file, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			panic(err)
		}

		defer file.Close()

		// Create a new slog logger with the log file
		logHandler := slog.NewTextHandler(file, &slog.HandlerOptions{
			Level:     slog.LevelDebug,
			AddSource: true,
		})

		logger := slog.New(logHandler)
		slog.SetDefault(logger)

	} else {
		// If logging to file is not enabled, create a new tint logger
		tintHandler := tint.NewHandler(os.Stderr, &tint.Options{
			Level:     slog.LevelDebug,
			AddSource: true,
		})

		tintLogger := slog.New(tintHandler)
		slog.SetDefault(tintLogger)
	}
}
