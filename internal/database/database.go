package database

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	glogger "gorm.io/gorm/logger"

	"github.com/MESMUR/fixed-term-track-web-server/config"
	"github.com/MESMUR/fixed-term-track-web-server/internal/models"
	"github.com/MESMUR/fixed-term-track-web-server/pkg/logger"
)

func ConnectPostgres() *gorm.DB {
	logger.Log.Info("Initializing database connection")

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=%s",
		config.AppConfig.DBHost,
		config.AppConfig.DBUser,
		config.AppConfig.DBPassword,
		config.AppConfig.DBName,
		config.AppConfig.DBPort,
		config.AppConfig.DBTimezone,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: glogger.Default.LogMode(glogger.Info),
	})

	if err != nil {
		logger.Sugar.Fatalf("Failed to connect to database: %v", err)
	}

	err = db.AutoMigrate(&models.FixedTerm{}, &models.FixedTermReturn{}, &models.Event{})

	if err != nil {
		logger.Sugar.Fatalf("Failed to migrate: %v", err)
	}

	return db
}
