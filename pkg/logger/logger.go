package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

var (
	// Log is the global logger
	Log *zap.Logger

	Sugar *zap.SugaredLogger
)

// Initialize initializes the global logger
func Initialize() {
	var err error
	config := zap.NewProductionConfig()

	// Set custom configurations
	config.EncoderConfig.TimeKey = "timestamp"
	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	config.OutputPaths = []string{"stdout"}
	config.ErrorOutputPaths = []string{"stderr"}

	// Check if running in development mode
	if os.Getenv("GO_ENV") == "development" {
		config = zap.NewDevelopmentConfig()
	}

	// Build the logger
	Log, err = config.Build()
	if err != nil {
		panic(err)
	}

	// Create a SugaredLogger for convenience
	Sugar = Log.Sugar()

	// Replace the global zap logger with the custom one
	zap.ReplaceGlobals(Log)
}

// Sync flushes any buffered log entries
func Sync() {
	_ = Log.Sync()
	_ = Sugar.Sync()
}
