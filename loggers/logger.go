package loggers

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func InitLogger() (*zap.Logger, error) {
	config := zap.NewProductionConfig()
	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	config.OutputPaths = []string{"./loggers/logger.log"}
	
	return config.Build()
}
