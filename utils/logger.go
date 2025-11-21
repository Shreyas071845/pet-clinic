package utils

import (
	"io"
	"os"

	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
)

var Log = logrus.New()

// initializes a global structured logger with multiple log levels.
func InitLogger() {

	// Create a log rotation setup
	fileLogger := &lumberjack.Logger{
		Filename:   "app.log",
		MaxSize:    5,
		MaxBackups: 3,
		MaxAge:     28,
		Compress:   true,
	}

	// Combine outputs: write logs to both console and file
	multiWriter := io.MultiWriter(os.Stdout, fileLogger)

	Log.SetOutput(multiWriter)

	// Use structured JSON format with timestamps
	Log.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})

	Log.SetLevel(logrus.DebugLevel)

	// Test logs to confirm setup
	Log.Debug("Logger initialized in DEBUG mode (verbose output)")
	Log.Info("Logger initialized successfully")
	Log.Warn("Logger rotation is active")
}
