package main

import (
	"github.com/gin-gonic/gin"
	"go-api/handler"
	"go-api/model"
	"go-api/repository"
	"go-api/service"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
)

func main() {

	dsn := os.Getenv("DB_URL")
	db, errCon := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if errCon != nil {
		log.Fatal(errCon.Error())
	}

	if err := db.AutoMigrate(&model.User{}); err != nil {
		log.Fatal(errCon.Error())
		return
	}

	userRepo := repository.NewRepository(db)
	userService := service.NewService(userRepo)
	userHandler := handler.Userhandler(userService)
	router := gin.Default()

	router.GET("/users/:id", userHandler.GetUser)
	router.GET("/users", userHandler.GetUsers)
	router.POST("/users", userHandler.StoreUser)

	err := router.Run()
	if err != nil {
		return
	}
}
