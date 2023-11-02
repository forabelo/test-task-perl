package deps

import (
	"fmt"
	"github.com/bradfitz/gomemcache/memcache"
	"github.com/forabelo/test-task-perl/config"
	"github.com/forabelo/test-task-perl/internal/repository"
	"github.com/forabelo/test-task-perl/internal/service"
	"github.com/forabelo/test-task-perl/internal/service/cache"
)

type Dependencies struct {
	Config     *config.Config
	CountryUcs *service.CountryUseCase
}

func Init(cfg *config.Config) (*Dependencies, error) {
	fmt.Println("Initializing dependencies...")

	repo, err := repository.NewGeoIP2Repository(cfg.GeoIP2DatabasePath)
	if err != nil {
		return nil, err
	}

	memcacheClient := memcache.New(cfg.MemcacheHost)

	ucs := service.NewIpNetworkUseCase(
		repo,
		cache.NewCompanyMemcacheCacheService(memcacheClient),
	)

	return &Dependencies{
		Config:     cfg,
		CountryUcs: ucs,
	}, nil
}
