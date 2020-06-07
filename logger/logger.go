package logger

import (
	"os"

	"github.com/caioeverest/transactions-api/config"
	"github.com/sirupsen/logrus"
)

type Logger struct {
	logrus.Logger
}

var (
	baseLogger *Logger
	level      logrus.Level
)

func Start() {
	switch config.Get().ENV {
	case config.DEV:
		level = logrus.DebugLevel
	default:
		level = logrus.InfoLevel
	}

	baseLogger = &Logger{
		logrus.Logger{
			Out:          os.Stderr,
			Formatter:    new(logrus.TextFormatter),
			Hooks:        make(logrus.LevelHooks),
			Level:        level,
			ExitFunc:     os.Exit,
			ReportCaller: false,
		},
	}

	if level == logrus.DebugLevel {
		baseLogger.SetFormatter(&logrus.TextFormatter{})
	} else {
		baseLogger.SetFormatter(&logrus.JSONFormatter{})
	}
}

func Get() *Logger { return baseLogger }
