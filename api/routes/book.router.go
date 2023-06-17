package routes

import (
	"github.com/gin-gonic/gin"
	handler "go-api/api/book"
	"go-api/db/model"
	"go-api/db/repository"
	"go-api/db/service"
	"gorm.io/gorm"
	"log"
)

func SetupBookRouter(db *gorm.DB, router *gin.RouterGroup) {
	if err := db.AutoMigrate(&model.Book{}); err != nil {
		log.Fatal(err.Error())
		return
	}
	bookRepo := repository.NewBookRepository(db)
	bookServ := service.NewBookService(bookRepo)
	bookHandler := handler.BookHandler(bookServ)

	router.GET("/", bookHandler.GetBooks)

}
