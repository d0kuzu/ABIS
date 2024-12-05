package rest

import (
	"ABIS/api/infrastructure/controllers/book_controllers"
	"github.com/gin-gonic/gin"
)

func BookRoutes(router *gin.Engine) {
	productGroup := router.Group("")
	{
		productGroup.POST("/create_book", book_controllers.CreateBook)
		productGroup.POST("/create_books", book_controllers.CreateBooks)
	}
}
