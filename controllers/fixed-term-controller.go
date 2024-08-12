package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/MESMUR/fixed-term-track-web-server/internal/database/models"
	"github.com/MESMUR/fixed-term-track-web-server/services"
)

type FixedTermController struct {
	fixedTermService services.FixedTermService
}

func NewFixedTermController(fixedTermService services.FixedTermService) *FixedTermController {
	return &FixedTermController{fixedTermService}
}

func (c *FixedTermController) GetFixedTermByID(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("fixed_term_id"), 10, 32)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid fixed term ID"})
		return
	}

	fixedTerm, err := c.fixedTermService.FindByID(uint(id))

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, fixedTerm)
}

func (c *FixedTermController) CreateFixedTerm(ctx *gin.Context) {
	var fixedTerm models.FixedTerm
	if err := ctx.ShouldBindJSON(&fixedTerm); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdFixedTerm, err := c.fixedTermService.Create(&fixedTerm)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, createdFixedTerm)
}

func (c *FixedTermController) UpdateFixedTerm(ctx *gin.Context) {
	var fixedTerm models.FixedTerm
	if err := ctx.ShouldBindJSON(&fixedTerm); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	updatedTerm, err := c.fixedTermService.Update(&fixedTerm)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, updatedTerm)
}

func (c *FixedTermController) CreateReturn(ctx *gin.Context) {
	var fixedTermReturn models.FixedTermReturn

	fixedTermID, err := strconv.ParseUint(ctx.Param("fixed_term_id"), 10, 32)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid fixed term ID"})
		return
	}

	fixedTermReturn.FixedTermID = uint(fixedTermID)

	if err = ctx.ShouldBindJSON(&fixedTermReturn); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdFixedTerm, err := c.fixedTermService.CreateReturn(&fixedTermReturn)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, createdFixedTerm)
}

func (c *FixedTermController) GetReturnByID(ctx *gin.Context) {
	fixedTermID, err := strconv.ParseUint(ctx.Param("fixed_term_id"), 10, 32)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid fixed term ID"})
		return
	}

	id, err := strconv.ParseUint(ctx.Param("return_id"), 10, 32)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid return ID"})
		return
	}

	fixedTermReturn, err := c.fixedTermService.FindReturnByID(uint(fixedTermID), uint(id))

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, fixedTermReturn)
}
