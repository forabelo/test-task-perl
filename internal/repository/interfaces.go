package repository

import (
	"github.com/forabelo/test-task-perl/internal/model"
	"net"
)

type IpCountryRepositoryInterface interface {
	GetCountryByIP(ip net.IP) (model.Country, error)
}
