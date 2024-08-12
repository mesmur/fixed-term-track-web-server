package routes

import (
	"github.com/MESMUR/fixed-term-track-web-server/controllers"
	"github.com/MESMUR/fixed-term-track-web-server/middleware"
	"github.com/MESMUR/fixed-term-track-web-server/pkg/logger"
	"github.com/gin-gonic/gin"
)

func SetupRouter(fixedTermController *controllers.FixedTermController) *gin.Engine {
	logger.Log.Info("Setting up routes")

	router := gin.New()

	router.Use(gin.Recovery())

	router.Use(middleware.BasicAuthentication())
	router.Use(middleware.RequestLogger())

	fixedTermRoutes := router.Group("/fixed_terms")
	{
		fixedTermRoutes.GET("/:fixed_term_id", fixedTermController.GetFixedTermByID)
		fixedTermRoutes.POST("/", fixedTermController.CreateFixedTerm)
		fixedTermRoutes.PUT("/", fixedTermController.UpdateFixedTerm)
		fixedTermRoutes.GET("/:fixed_term_id/returns/:return_id", fixedTermController.GetReturnByID)
		fixedTermRoutes.POST("/:fixed_term_id/returns", fixedTermController.CreateReturn)
	}

	return router
}
