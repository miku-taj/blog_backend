package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"project/internal/handler"
	"project/internal/repository"
	"project/internal/service"
)

func main() {

	fmt.Println("Library System")

	db, err := connectToDB()
	if err != nil {
		fmt.Printf("Failed to connect to Database: %v\n", err)
		return
	}

	blogRepository := repository.NewRepository(db)
	blogService := service.NewService(blogRepository)

	router := gin.Default()
	blogHandler := handler.NewHandler(router, blogService)
	blogHandler.InitRoutes()

	router.Run(":8080")
	log.Fatalf("Failed to start server: %v", router.Run(":8080"))

	els, _ := (blogRepository.GetRecentPostComments(1, 1))
	for el := range els {
		fmt.Println(*els[el])
	}
}

func connectToDB() (*gorm.DB, error) {
	dsn := "host=localhost user=muqaddas password=password dbname=blog_db port=5432 sslmode=disable TimeZone=Asia/Dushanbe"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect db: %v\n", err)
	}
	return db, nil
}
