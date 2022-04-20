package di

import (
	"tix-rest-project/internal/shared/config"

	"github.com/digitalysin/goblog/logger"
)

func NewLogger(cfg *config.Configuration) (logger.Logger, error) {
	return logger.New(&logger.Option{
		Level: logger.Info,
	})
}
