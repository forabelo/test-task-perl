package service

import (
	"context"
	"errors"
	"github.com/forabelo/test-task-perl/internal/model"
	"github.com/forabelo/test-task-perl/internal/repository"
	"github.com/forabelo/test-task-perl/internal/service/cache"
	"net"
)

type CountryUseCase struct {
	repo  repository.IpCountryRepositoryInterface
	cache cache.CompanyCacheServiceInterface
}

func NewIpNetworkUseCase(
	repo repository.IpCountryRepositoryInterface, cache cache.CompanyCacheServiceInterface) *CountryUseCase {
	return &CountryUseCase{
		repo:  repo,
		cache: cache,
	}
}

func (r *CountryUseCase) GetCountryByIp(ctx context.Context, ip string) (model.Country, error) {
	// Parsing ip address
	ipAddr := net.ParseIP(ip)
	if ipAddr == nil {
		return model.Country{}, errors.New("invalid ip address")
	}

	// Try to cache
	country, err := r.cache.GetCountryByIP(ip)
	if err == nil {
		return country, nil
	}

	// Get country from repository
	country, err = r.repo.GetCountryByIP(ipAddr)
	if err != nil {
		return model.Country{}, err
	}

	// Save in the cache
	_ = r.cache.SetCountryByIP(ip, country)

	return country, nil
}
