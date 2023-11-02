package repository

import (
	"fmt"
	"github.com/forabelo/test-task-perl/internal/model"
	"github.com/oschwald/geoip2-golang"
	"net"
)

type GeoIP2Repository struct {
	db *geoip2.Reader
}

func NewGeoIP2Repository(path string) (*GeoIP2Repository, error) {
	db, err := geoip2.Open(path)
	if err != nil {
		return nil, err
	}

	return &GeoIP2Repository{db: db}, nil
}

func (r *GeoIP2Repository) GetCountryByIP(ip net.IP) (model.Country, error) {
	fmt.Println("Request in GeoIP2")

	gIP2Country, err := r.db.Country(ip)
	if err != nil {
		return model.Country{}, err
	}

	return model.Country{
		Ip:          ip.String(),
		ISOCode:     gIP2Country.Country.IsoCode,
		CountryName: gIP2Country.Country.Names["en"],
	}, nil
}

func (r *GeoIP2Repository) Close() error {
	return r.db.Close()
}
