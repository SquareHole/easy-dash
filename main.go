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
	"github.com/squarehole/easydash/api"
	"github.com/squarehole/easydash/data"
	"github.com/squarehole/easydash/scheduling"
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

		err := data.InitDatabase()
		if err != nil {
			slog.Error("Error while initializing the database", "error", err.Error())
			panic(err)
		}
	}()

	// Create a new scheduler
	rs := scheduling.New()

	// Start the web server =, passing in the scehduler
	log.Fatal(api.Serve(rs))
}

func setupLogger() {

	logToFile, err := strconv.ParseBool(os.Getenv("LOG_TO_FILE"))
	if err != nil {
		logToFile = false
	}

	if logToFile {

		fileName := os.Getenv("LOG_FILE")
		fmt.Println(fileName)
		file, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			panic(err)
		}

		defer file.Close()

		logHandler := slog.NewTextHandler(file, &slog.HandlerOptions{
			Level:     slog.LevelDebug,
			AddSource: true,
		})

		logger := slog.New(logHandler)
		slog.SetDefault(logger)

	} else {
		tintHandler := tint.NewHandler(os.Stderr, &tint.Options{
			Level:     slog.LevelDebug,
			AddSource: true,
		})

		tintLogger := slog.New(tintHandler)
		slog.SetDefault(tintLogger)
	}
}
