package controllers

import (
	"net/http"
	"strconv"

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
	metric, err := c.service.GetTotalInvestedToDate()

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Unable to get total invested to date"})
		return
	}

	ctx.JSON(http.StatusOK, metric)
}

func (c *MetricsController) GetTotalCurrentlyInvested(ctx *gin.Context) {
	metric, err := c.service.GetTotalCurrentlyInvested()

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Unable to get total currently invested"})
		return
	}

	ctx.JSON(http.StatusOK, metric)
}

func (c *MetricsController) GetTotalMaturingInMonths(ctx *gin.Context) {
	monthsStr := ctx.Query("months")
	months, err := strconv.Atoi(monthsStr)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "months must be an integer"})
		return
	}

	metric, err := c.service.GetTotalMaturingInMonths(months)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Unable to get total maturing in months"})
		return
	}

	ctx.JSON(http.StatusOK, metric)
}

func (c *MetricsController) GetTotalReturnsToDate(ctx *gin.Context) {
	metric, err := c.service.GetTotalReturnsToDate()

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Unable to get total returns to date"})
		return
	}

	ctx.JSON(http.StatusOK, metric)
}

func (c *MetricsController) GetTotalReturnsThisYear(ctx *gin.Context) {
	metric, err := c.service.GetTotalReturnsThisYear()

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Unable to get total returns this year"})
		return
	}

	ctx.JSON(http.StatusOK, metric)
}
