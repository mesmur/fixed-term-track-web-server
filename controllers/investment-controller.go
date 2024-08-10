package controllers

import (
	"github.com/MESMUR/fixed-term-track-web-server/internal/database/models"
	"github.com/MESMUR/fixed-term-track-web-server/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type InvestmentController struct {
	investmentService services.InvestmentService
}

func NewUserController(userService services.InvestmentService) *InvestmentController {
	return &InvestmentController{userService}
}

func (c *InvestmentController) GetInvestmentByID(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("investment_id"), 10, 32)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	user, err := c.investmentService.FindByID(uint(id))

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, user)
}

func (c *InvestmentController) CreateInvestment(ctx *gin.Context) {
	var investment models.Investment
	if err := ctx.ShouldBindJSON(&investment); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdUser, err := c.investmentService.Create(&investment)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, createdUser)
}

func (c *InvestmentController) UpdateInvestment(ctx *gin.Context) {
	var investment models.Investment
	if err := ctx.ShouldBindJSON(&investment); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	updatedInvestment, err := c.investmentService.Update(&investment)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, updatedInvestment)
}

func (c *InvestmentController) CreateInvestmentReturn(ctx *gin.Context) {
	var investmentReturn models.InvestmentReturn

	investmentId, err := strconv.ParseUint(ctx.Param("investment_id"), 10, 32)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid investment ID"})
		return
	}

	investmentReturn.InvestmentID = uint(investmentId)

	if err = ctx.ShouldBindJSON(&investmentReturn); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdInvestmentReturn, err := c.investmentService.CreateInvestmentReturn(&investmentReturn)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, createdInvestmentReturn)
}

func (c *InvestmentController) GetInvestmentReturnByID(ctx *gin.Context) {
	investmentId, err := strconv.ParseUint(ctx.Param("investment_id"), 10, 32)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid investment ID"})
		return
	}

	id, err := strconv.ParseUint(ctx.Param("return_id"), 10, 32)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid return ID"})
		return
	}

	investmentReturn, err := c.investmentService.FindReturnByID(uint(investmentId), uint(id))

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, investmentReturn)
}
