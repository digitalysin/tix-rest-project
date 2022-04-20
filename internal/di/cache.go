package di

import (
	"tix-rest-project/internal/shared/config"

	"github.com/digitalysin/goblog/cache"
)

func NewCache(cfg *config.Configuration) (cache.Cache, error) {
	return cache.New(&cache.Option{
		Address: cfg.CacheHost,
	})
}
