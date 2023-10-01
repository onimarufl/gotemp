package main

import (
	"fmt"
	"gotemp/handler"
	"gotemp/repository"
	"gotemp/service"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func main() {

	db, err := OpenDatabase()
	if err != nil {
		panic(err)
	}

	r := gin.Default()

	repository := repository.NewRepository(db)
	service := service.NewService(repository)
	handler := handler.NewHandler(service)

	r.GET("/getMessage", handler.HandlerGetMessage)
	r.POST("/postService", handler.HandlerPostService)

	r.POST("/user", handler.HandlerInsertData)
	r.GET("/user", handler.HandlerInquiryDataByID)
	r.GET("/user/all", handler.HandlerInquiryAllData)
	r.PUT("/user", handler.HandlerUpdateDataByID)
	r.DELETE("/user", handler.HandlerDeleteDataByID)

	r.Run()
}

func OpenDatabase() (*sqlx.DB, error) {
	var dataSoruce = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s", "localhost", 5432, "postgres", "root1234", "go_test", "disable")
	return sqlx.Connect("postgres", dataSoruce)
}
