package main

import (
	"fmt"
	"log"
	"os"

	"gits/auth"
	"gits/author"
	"gits/book"
	"gits/handler"
	Middleware "gits/middleware"
	"gits/publisher"
	"gits/redisClient"
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
		&user.User{},
		&author.Author{},
		&publisher.Publisher{},
		&book.Book{},
	)

	fmt.Println("\n koneksi dataBase berhasil *******\n")

	var redisHost = "localhost:6379"
	var redisPassword = "root"

	rds := redisClient.ConnectRedisClient(redisHost, redisPassword)

	fmt.Println("\n koneksi Redis berhasil *******\n")

	userRepository := user.NewRepository(dataBase)
	authorRepository := author.NewRepository(dataBase, rds)
	publisherRepository := publisher.NewRepository(dataBase, rds)
	bookRepository := book.NewRepository(dataBase, rds)

	userService := user.NewService(userRepository)
	authorService := author.NewService(authorRepository)
	publisherService := publisher.NewService(publisherRepository)
	bookService := book.NewService(bookRepository)
	authService := auth.NewService()

	userHandler := handler.NewUserHandler(userService, authService)
	authorHandler := handler.NewAuthorHandler(authorService)
	publisherHandler := handler.NewPublisherHandler(publisherService)
	bookHandler := handler.NewBookHandler(bookService)

	router := gin.Default()
	api := router.Group("/api")

	api.POST("/user/register", userHandler.RegisterUser)
	api.POST("/user/login", userHandler.Login)

	api.POST("/author/register", Middleware.IsLogin(authService, userService), authorHandler.RegisterAuthor)
	api.GET("/author/:id", Middleware.IsLogin(authService, userService), authorHandler.GetAuthorById)
	api.PUT("/author/:id", Middleware.IsLogin(authService, userService), authorHandler.UpdateAuthorById)
	api.DELETE("/author/:id", Middleware.IsLogin(authService, userService), authorHandler.DestroyAuthorById)

	api.POST("/publisher/register", Middleware.IsLogin(authService, userService), publisherHandler.RegisterPublisher)
	api.GET("/publisher/:id", Middleware.IsLogin(authService, userService), publisherHandler.GetPublisherById)
	api.PUT("/publisher/:id", Middleware.IsLogin(authService, userService), publisherHandler.UpdatePublisherById)
	api.DELETE("/publisher/:id", Middleware.IsLogin(authService, userService), publisherHandler.DestroyPublisherById)

	api.POST("/book/register", Middleware.IsLogin(authService, userService), bookHandler.RegisterBook)
	api.GET("/book/:id", Middleware.IsLogin(authService, userService), bookHandler.GetBookById)
	api.PUT("/book/:id", Middleware.IsLogin(authService, userService), bookHandler.UpdateBookById)
	api.DELETE("/book/:id", Middleware.IsLogin(authService, userService), bookHandler.DestroyBookById)

	router.Run(":3000")

}
