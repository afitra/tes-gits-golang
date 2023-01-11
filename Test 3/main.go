package main

import (
	"fmt"
	"log"
	"os"

	"gits/auth"
	"gits/handler"
	"gits/user"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file, please check your .env file")
	}

	var DB_PORT = os.Getenv("DB_PORT")
	if os.Getenv("DB_PORT") == "" {
		DB_PORT = "5432"
	}

	var PG_USERNAME = []byte(os.Getenv("PG_USERNAME"))
	var PG_PASSWORD = []byte(os.Getenv("PG_PASSWORD"))
	var PG_HOST = []byte(os.Getenv("PG_HOST"))
	var DB_NAME = []byte(os.Getenv("DB_NAME"))

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable ",
		PG_HOST, PG_USERNAME, PG_PASSWORD, DB_NAME, DB_PORT)
	dataBase, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	dataBase.Debug().AutoMigrate(
		&user.User{})
	fmt.Println("\n koneksi dataBase berhasil *******\n")

	userRepository := user.NewRepository(dataBase)

	userService := user.NewService(userRepository)
	authService := auth.NewService()

	userHandler := handler.NewUserHandler(userService, authService)

	router := gin.Default()
	api := router.Group("/api")

	api.POST("/register", userHandler.RegisterUser)
	api.POST("/login", userHandler.Login)

	router.Run(":3000")

}
