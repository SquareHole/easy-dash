package main

import (
	"log"
	"sync"

	"log/slog"

	"github.com/joho/godotenv"
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

// https://localhost:5341

func main() {

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()

		err := data.InitDatabase()
		if err != nil {
			slog.Error("Error while initializing the database", "error", err.Error())
			panic(err)
		}
	}()

	rs := scheduling.Scehdule()
	log.Fatal(api.Serve(rs))
}
