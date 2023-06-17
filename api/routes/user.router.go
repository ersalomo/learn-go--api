package routes

import (
	"github.com/gin-gonic/gin"
	"go-api/api/book"
	"go-api/api/middleware"
	"go-api/db/model"
	"go-api/db/repository"
	"go-api/db/service"
	"gorm.io/gorm"
	"log"
)

func SetupUserRoute(db *gorm.DB, router *gin.RouterGroup) {
	if err := db.AutoMigrate(&model.User{}); err != nil {
		log.Fatal(err.Error())
		return
	}
	userRepo := repository.NewRepository(db)
	userService := service.NewService(userRepo)
	userHandler := handler.UserHandler(userService)

	router.POST("/", userHandler.StoreUser)
	protected := router.Use(middleware.DeserializeMiddleware())
	protected.GET("/", userHandler.GetUsers)
	protected.GET("/:id", userHandler.GetUser)
}
