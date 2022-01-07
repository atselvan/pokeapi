package pokeapi

import (
	"github.com/allegro/bigcache"
	"github.com/privatesquare/bkst-go-utils/utils/logger"
	"time"
)

const (
	defaultCacheTimeout = 10 * time.Minute
	defaultCleanWindow  = 5 * time.Minute
)

var (
	cache *bigcache.BigCache
	err   error
)

func init() {
	config := bigcache.Config{
		LifeWindow:  defaultCacheTimeout,
		CleanWindow: defaultCleanWindow,
	}
	if cache, err = bigcache.NewBigCache(config); err != nil {
		logger.Panic(err.Error())
	}
}
