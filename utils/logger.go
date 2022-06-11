package utils

import (
	"fmt"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func NewLogger() (*zap.Logger, error) {
	ENV := os.Getenv("ENV")

	var consoleEncoder zapcore.Encoder

	if ENV == "development" {
		consoleEncoder = zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())
	} else if ENV == "production" {
		consoleEncoder = zapcore.NewConsoleEncoder(zap.NewProductionEncoderConfig())
	} else {
		panic(fmt.Errorf("unknown env"))
	}

	infoLevel := zap.LevelEnablerFunc(func(l zapcore.Level) bool {
		return l == zapcore.InfoLevel
	})

	errorAndFatalLevel := zap.LevelEnablerFunc(func(l zapcore.Level) bool {
		return l == zapcore.ErrorLevel || l == zapcore.FatalLevel
	})

	core := zapcore.NewTee(
		zapcore.NewCore(
			consoleEncoder,
			zapcore.Lock(os.Stdout),
			infoLevel,
		),
		zapcore.NewCore(
			consoleEncoder,
			zapcore.Lock(os.Stderr),
			errorAndFatalLevel,
		),
	)

	logger := zap.New(core, zap.AddCaller())
	return logger, nil
}
