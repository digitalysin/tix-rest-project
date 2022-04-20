package di

import (
	"time"
	"tix-rest-project/internal/shared/config"

	"github.com/digitalysin/goblog/db"
	"github.com/digitalysin/goblog/logger"
)

func NewDatabase(cfg *config.Configuration, logger logger.Logger) (db.ORM, error) {
	return db.NewMySql(&db.MySqlOption{
		ConnectionString:      cfg.DatabaseConnectionString,
		MaxLifeTimeConnection: time.Hour,
		MaxIdleConnection:     8,
		MaxOpenConnection:     32,
		Logger:                logger,
	})
}
