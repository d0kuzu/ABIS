package api

import (
	"ABIS/api/infrastructure/rest"
	"github.com/gin-gonic/gin"
	"log"
)

func RouterStart() {
	r := gin.Default()

	rest.BookRoutes(r)

	//// Определение маршрута для GET-запроса на корень
	//r.GET("/", func(c *gin.Context) {
	//	c.JSON(http.StatusOK, gin.H{"message": "Hello, World!"})
	//})

	// Определение маршрута для GET-запроса на /ping
	//r.GET("/ping", func(c *gin.Context) {
	//	c.JSON(http.StatusOK, gin.H{"message": "pong"})
	//})
	//
	//r.POST("/test", func(c *gin.Context) {
	//
	//	user_repo.CreateUser(c)
	//})
	//
	//r.POST("/user_model", func(c *gin.Context) {
	//	var user_model struct {
	//		Name  string `json:"name"`
	//		Email string `json:"email"`
	//	}
	//
	//	// Привязка JSON-данных к структуре user_model
	//	if err := c.ShouldBindJSON(&user_model); err != nil {
	//		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	//		return
	//	}
	//
	//	// Возврат успешного ответа
	//	c.JSON(http.StatusOK, gin.H{"status": "user_model created", "user_model": user_model})
	//})

	err := r.Run(":8080")
	if err != nil {
		log.Fatal("Router start error", err)
	}
}
