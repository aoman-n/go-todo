package main

import (
	"github.com/gin-gonic/gin"
	"github.com/laster18/go-todo/config"
	"github.com/laster18/go-todo/controller"
)

func main() {
	router := gin.Default()

	router.GET("/", controller.IndexGET)
	router.POST("/", controller.PostsPost)
	router.Run(":" + config.Config.Port)
}
