package data

import (
	"os"

	"log/slog"

	"github.com/squarehole/easydash/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB
var err error

func InitDatabase() error {
	dsn := os.Getenv("CONNECTION")
	slog.Info("Initializing database", "connection", dsn)

	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		slog.Error(err.Error())
		return err
	}

	err = db.AutoMigrate(&models.UrlConfig{}, &models.Result{}, &models.Summary{}, &models.TestStatus{})
	if err != nil {
		slog.Error(err.Error())
		return err
	}

	return nil
}

func GetAllConfigs() ([]models.UrlConfig, error) {
	var configs []models.UrlConfig
	err := db.Find(&configs).Error
	if err != nil {
		slog.Error(err.Error())
		return nil, err
	}

	slog.Info("GetAllConfigs", "count", len(configs))
	return configs, nil
}
