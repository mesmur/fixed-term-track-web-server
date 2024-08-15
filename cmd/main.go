package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

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

	/**
	 * Graceful shutdown logic taken from the Gin documentation
	 * https://github.com/gin-gonic/examples/blob/master/graceful-shutdown/graceful-shutdown/notify-without-context/server.go
	 */
	srv := &http.Server{
		Addr:    config.AppConfig.Port,
		Handler: router,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("Failed to listen to Addr: %s\n", err)
		}
	}()

	quit := make(chan os.Signal, 1)

	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	logger.Sugar.Info("Shutting down server...")

	// The context is used to inform the server it has 5 seconds to finish the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		logger.Sugar.Fatal("Server forced to shut down: ", err)
	}

	logger.Log.Info("Server exiting")
}

func init() {
	gin.SetMode(gin.ReleaseMode)
}
