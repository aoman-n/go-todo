package main

import (
	"github.com/gin-gonic/gin"
	"github.com/laster18/go-todo/config"
	"github.com/laster18/go-todo/controller"
)

func main() {
	router := gin.Default()

	v1 := router.Group("/api/v1")
	{
		v1.GET("/", controller.TodosIndexEndpoint)
		v1.GET("/:id", controller.TodosShowEndpoint)
		v1.POST("/", controller.TodosCreateEndpoint)
		v1.DELETE("/:id", controller.TodosDeleteEndpoint)
	}

	router.Run(":" + config.Config.Port)
}
