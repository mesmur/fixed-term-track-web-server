package config

import (
	"github.com/MESMUR/fixed-term-track-web-server/pkg/logger"
	"github.com/spf13/viper"
	"log"
	"time"
)

type Config struct {
	DBHost           string
	DBPort           string
	DBUser           string
	DBPassword       string
	DBName           string
	DBTimezone       string
	Port             string
	AppUsername      string
	AppPassword      string
	TelegramBotToken string
	TelegramChatID   int64
}

var AppConfig *Config

func LoadConfig() {
	logger.Log.Info("Loading config")

	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	AppConfig = &Config{
		DBHost:           viper.GetString("DB_HOST"),
		DBPort:           viper.GetString("DB_PORT"),
		DBUser:           viper.GetString("DB_USER"),
		DBPassword:       viper.GetString("DB_PASSWORD"),
		DBName:           viper.GetString("DB_NAME"),
		DBTimezone:       viper.GetString("DB_TIMEZONE"),
		Port:             viper.GetString("PORT"),
		AppUsername:      viper.GetString("APP_USERNAME"),
		AppPassword:      viper.GetString("APP_PASSWORD"),
		TelegramBotToken: viper.GetString("TELEGRAM_BOT_TOKEN"),
		TelegramChatID:   viper.GetInt64("TELEGRAM_CHAT_ID"),
	}

	time.Local = time.UTC
}
