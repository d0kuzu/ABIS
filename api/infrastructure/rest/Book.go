package rest

import (
	"ABIS/api/infrastructure/controllers/book_controllers"
	"github.com/gin-gonic/gin"
)

func BookRoutes(router *gin.Engine) {
	productGroup := router.Group("/phones")
	{
		productGroup.POST("/", book_controllers.Create)
	}
}
