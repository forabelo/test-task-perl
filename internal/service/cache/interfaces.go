package cache

import "github.com/forabelo/test-task-perl/internal/model"

// We could use additional interface for cache service (e.g., cache interface), but it's not necessary for this task

type CompanyCacheServiceInterface interface {
	GetCountryByIP(ip string) (model.Country, error)
	SetCountryByIP(ip string, country model.Country) error
}
