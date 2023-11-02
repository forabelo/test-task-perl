package handlers

import (
	"context"
	"github.com/forabelo/test-task-perl/internal/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CountryUseCaseInterface interface {
	GetCountryByIp(ctx context.Context, ip string) (model.Country, error)
}

type CountryHandler struct {
	ucs CountryUseCaseInterface
}

func NewCountryHandler(ucs CountryUseCaseInterface) *CountryHandler {
	return &CountryHandler{
		ucs: ucs,
	}
}

func (h *CountryHandler) GetCountryByIp(c *gin.Context) {
	ip := c.Query("ip")

	if ip == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "ip is required",
		})
		return
	}

	ctx := c.Request.Context()

	e, err := h.ucs.GetCountryByIp(ctx, ip)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Can be additional formatting into a response model. But it's not required by task

	c.JSON(http.StatusOK, e)
}
