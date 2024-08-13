package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/MESMUR/fixed-term-track-web-server/services"
)

type MetricsController struct {
	service services.MetricsService
}

func NewMetricsController(service services.MetricsService) *MetricsController {
	return &MetricsController{service}
}

func (c *MetricsController) GetTotalInvestedToDate(ctx *gin.Context) {
	var total float64
	total, err := c.service.GetTotalInvestedToDate()

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Unable to get total invested to date"})
		return
	}

	ctx.JSON(http.StatusOK, total)
}

func (c *MetricsController) GetTotalCurrentlyInvested(ctx *gin.Context) {
	var total float64
	total, err := c.service.GetTotalCurrentlyInvested()

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Unable to get total currently invested"})
		return
	}

	ctx.JSON(http.StatusOK, total)
}
