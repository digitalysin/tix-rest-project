package shared

import (
	"tix-rest-project/internal/shared/config"

	"github.com/digitalysin/goblog/cache"
	"github.com/digitalysin/goblog/db"
	"github.com/digitalysin/goblog/logger"
	"github.com/labstack/echo/v4"
	"go.uber.org/dig"
)

type (
	Deps struct {
		dig.In
		Config   *config.Configuration
		Logger   logger.Logger
		Database db.ORM
		Cache    cache.Cache
		Echo     *echo.Echo
	}
)

func (impl *Deps) Close() {
	if err := impl.Database.Close(); err != nil {
		impl.Logger.Errorf("failed to close db: %s", err)
	}

	if err := impl.Cache.Close(); err != nil {
		impl.Logger.Errorf("failed to close cache: %s", err)
	}

	if err := impl.Echo.Close(); err != nil {
		impl.Logger.Errorf("failed to http server: %s", err)
	}
}

func Register(container *dig.Container) error {
	if err := container.Provide(config.New); err != nil {
		return err
	}

	return nil
}
