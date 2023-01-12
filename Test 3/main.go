package main

import (
	"fmt"
	"log"
	"os"

	"gits/auth"
	"gits/author"
	"gits/book"
	"gits/handler"
	"gits/publisher"
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

	userRepository := user.NewRepository(dataBase)
	authorRepository := author.NewRepository(dataBase)
	publisherRepository := publisher.NewRepository(dataBase)
	bookRepository := book.NewRepository(dataBase)

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

	api.POST("/author/register", authorHandler.RegisterAuthor)
	api.GET("/author/:id", authorHandler.GetAuthorById)
	api.PUT("/author/:id", authorHandler.UpdateAuthorById)
	api.DELETE("/author/:id", authorHandler.DestroyAuthorById)

	api.POST("/publisher/register", publisherHandler.RegisterPublisher)
	api.GET("/publisher/:id", publisherHandler.GetPublisherById)
	api.PUT("/publisher/:id", publisherHandler.UpdatePublisherById)
	api.DELETE("/publisher/:id", publisherHandler.DestroyPublisherById)

	api.POST("/book/register", bookHandler.RegisterBook)
	api.GET("/book/:id", bookHandler.GetBookById)
	api.PUT("/book/:id", bookHandler.UpdateBookById)
	api.DELETE("/book/:id", bookHandler.DestroyBookById)

	router.Run(":3000")

}
