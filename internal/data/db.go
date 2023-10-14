package data

import (
	"os"

	"log/slog"

	"github.com/squarehole/easydash/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// db is a global variable that holds the database connection
var db *gorm.DB
var err error

// InitDatabase initializes the database connection and performs migrations
func InitDatabase() error {
	// Get the database connection string from the environment
	dsn := os.Getenv("CONNECTION")

	// Log the database initialization
	slog.Info("Initializing database", "connection", dsn)

	// Open a connection to the database
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		// Log the error if the connection fails
		slog.Error(err.Error())
		return err
	}

	// Perform database migrations
	err = db.AutoMigrate(&models.UrlConfig{}, &models.Result{}, &models.Summary{}, &models.TestStatus{})
	if err != nil {
		// Log the error if migrations fail
		slog.Error(err.Error())
		return err
	}

	return nil
}

// GetAllConfigs retrieves all UrlConfig objects from the database
func GetAllConfigs() ([]models.UrlConfig, error) {
	var configs []models.UrlConfig

	// Query the database for all UrlConfig objects
	err := db.Find(&configs).Error
	if err != nil {
		// Log the error if the query fails
		slog.Error(err.Error())
		return nil, err
	}

	// Log the number of UrlConfig objects retrieved
	slog.Info("GetAllConfigs", "count", len(configs))
	return configs, nil
}
