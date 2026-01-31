package util

import (
	"fmt"

	"go.uber.org/zap"
)

func InitLogger(env string) (*zap.SugaredLogger, error) {
	var (
		logger *zap.Logger
		err    error
	)
	switch env {
	case "dev":
		logger, err = zap.NewDevelopment()
	case "prod":
		logger, err = zap.NewProduction()
	default:
		return nil, fmt.Errorf("Не известное окружение %s", env)
	}

	return logger.Sugar(), err
}
