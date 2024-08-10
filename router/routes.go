package routes

import (
	"github.com/MESMUR/fixed-term-track-web-server/controllers"
	"github.com/MESMUR/fixed-term-track-web-server/middleware"
	"github.com/MESMUR/fixed-term-track-web-server/pkg/logger"
	"github.com/gin-gonic/gin"
)

func SetupRouter(investmentController *controllers.InvestmentController) *gin.Engine {
	logger.Log.Info("Setting up routes")

	router := gin.New()

	router.Use(gin.Recovery())

	router.Use(middleware.BasicAuthentication())
	router.Use(middleware.RequestLogger())

	investmentRoutes := router.Group("/investments")
	{
		investmentRoutes.GET("/:investment_id", investmentController.GetInvestmentByID)
		investmentRoutes.POST("/", investmentController.CreateInvestment)
		investmentRoutes.PUT("/", investmentController.UpdateInvestment)
		investmentRoutes.GET("/:investment_id/returns/:return_id", investmentController.GetInvestmentReturnByID)
		investmentRoutes.POST("/:investment_id/returns", investmentController.CreateInvestmentReturn)
	}

	return router
}
