package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go-api/api/routes"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
)

func main() {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error load .env file")
	}
	dsn := os.Getenv("DB_URL")

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	router := gin.Default()

	userRouter := router.Group("/api/users")
	routes.SetupUserRoute(db, userRouter)

	bookRouter := router.Group("/api/books")
	routes.SetupBookRouter(db, bookRouter)

	err = router.Run(":8008")

	if err != nil {
		return
	}
}
