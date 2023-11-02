package cache

import (
	"encoding/json"
	"github.com/bradfitz/gomemcache/memcache"
	"github.com/forabelo/test-task-perl/internal/model"
)

type CompanyMemcacheCacheService struct {
	client *memcache.Client
}

func NewCompanyMemcacheCacheService(client *memcache.Client) *CompanyMemcacheCacheService {
	return &CompanyMemcacheCacheService{client: client}
}

func (s *CompanyMemcacheCacheService) GetCountryByIP(ip string) (model.Country, error) {
	item, err := s.client.Get(ip)
	if err == nil && item != nil {
		var cachedCountry model.Country
		err := json.Unmarshal(item.Value, &cachedCountry)
		if err == nil {
			return cachedCountry, nil
		}
	}

	return model.Country{}, err
}

func (s *CompanyMemcacheCacheService) SetCountryByIP(ip string, country model.Country) error {
	data, err := json.Marshal(country)
	if err == nil {
		return s.client.Set(&memcache.Item{Key: ip, Value: data, Expiration: 3600}) // 1 hours
	}

	return err
}
