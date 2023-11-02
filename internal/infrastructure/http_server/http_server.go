package http_server

import (
	"github.com/forabelo/test-task-perl/internal/app/deps"
	handlers "github.com/forabelo/test-task-perl/internal/controller/http/v1/httphandlers"
	"github.com/gin-gonic/gin"
)

type HttpServer interface {
	Run(addr ...string) error
}

func NewHTTPServer(d *deps.Dependencies) (HttpServer, error) {
	if isProduction(d.Config.AppEnv) {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}

	r := gin.Default()

	err := r.SetTrustedProxies(nil)
	if err != nil {
		return nil, err
	}

	h := handlers.NewCountryHandler(d.CountryUcs)

	r.GET("/api/v1/ip-networks/country", h.GetCountryByIp)

	return r, nil
}

func isProduction(appEnv string) bool {
	return appEnv == "prod"
}
