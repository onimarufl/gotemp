package main

import (
	"gotemp/handler"
	"gotemp/service"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	service := service.NewService()
	handler := handler.NewHandler(service)

	r.GET("/getMessage", handler.HandlerGetMessage)
	r.POST("/postService", handler.HandlerPostService)

	r.Run()
}
