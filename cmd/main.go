package main

import (
	"github.com/gin-gonic/gin"

	"github.com/MESMUR/fixed-term-track-web-server/config"
	"github.com/MESMUR/fixed-term-track-web-server/controllers"
	"github.com/MESMUR/fixed-term-track-web-server/cron"
	"github.com/MESMUR/fixed-term-track-web-server/internal/clients"
	"github.com/MESMUR/fixed-term-track-web-server/internal/database"
	"github.com/MESMUR/fixed-term-track-web-server/pkg/logger"
	"github.com/MESMUR/fixed-term-track-web-server/repositories"
	routes "github.com/MESMUR/fixed-term-track-web-server/router"
	"github.com/MESMUR/fixed-term-track-web-server/services"
)

func main() {
	logger.Initialize()
	defer logger.Sync()

	config.LoadConfig()

	db := database.ConnectPostgres()

	fixedTermRepository := repositories.NewFixedTermRepository(db)
	fixedTermReturnRepository := repositories.NewFixedTermReturnRepository(db)
	eventRepository := repositories.NewEventRepository(db)
	fixedTermService := services.NewFixedTermService(fixedTermRepository, fixedTermReturnRepository, eventRepository)
	fixedTermController := controllers.NewFixedTermController(fixedTermService)

	metricsRepository := repositories.NewMetricsRepository(db)
	metricsService := services.NewMetricsService(metricsRepository)
	metricsController := controllers.NewMetricsController(metricsService)

	router := routes.SetupRouter(fixedTermController, metricsController)

	telegramSdk := clients.CreateTelegramSdk(config.AppConfig.TelegramBotToken, config.AppConfig.TelegramChatID)
	eventReader := cron.NewEventReader(eventRepository, telegramSdk)
	go eventReader.CheckEvents()

	logger.Sugar.Infof("Starting server on port %s", config.AppConfig.Port)

	err := router.Run(config.AppConfig.Port)

	if err != nil {
		panic(err)
	}
}

func init() {
	gin.SetMode(gin.ReleaseMode)
}
