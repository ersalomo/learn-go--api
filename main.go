package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go-api/api/user_handler"
	"go-api/db/model"
	"go-api/db/repository"
	"go-api/db/service"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
)

func main() {

	//dsn := os.Getenv("DB_URL")
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error load .env file")
	}
	dsn := os.Getenv("DB_URL")

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	if err = db.AutoMigrate(&model.User{}); err != nil {
		log.Fatal(err.Error())
		return
	}

	userRepo := repository.NewRepository(db)
	userService := service.NewService(userRepo)
	userHandler := handler.Userhandler(userService)
	router := gin.Default()
	v1 := router.Group("/api/v1")

	v1.POST("/users", userHandler.StoreUser)

	protected := router.Group("/api/v1/authenticated")
	//protected.Use(middleware.DeserilizeMiddleware())
	protected.GET("/users", userHandler.GetUsers)
	protected.GET("/users/:id", userHandler.GetUser)

	err = router.Run(":8008")
	if err != nil {
		return
	}
}
