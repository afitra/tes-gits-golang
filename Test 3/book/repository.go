package book

import (
	"encoding/json"
	"fmt"
	"gits/redisClient"
	"strconv"

	"github.com/go-redis/redis/v9"
	"gorm.io/gorm"
)

type Repository interface {
	Create(book Book) (Book, error)
	FindById(ID int) (Book, error)
	Update(book Book) (Book, error)
	Destroy(ID int) error
}

type repository struct {
	db  *gorm.DB
	rds *redis.Client
}

func NewRepository(db *gorm.DB, rds *redis.Client) *repository {
	return &repository{db, rds}

}

func (r *repository) Create(book Book) (Book, error) {

	err := r.db.Create(&book).Error

	if err != nil {
		return book, err
	}
	return book, nil
}

func (r *repository) FindById(ID int) (Book, error) {

	key := "book" + strconv.Itoa(ID)
	jsonString := redisClient.GetData(r.rds, key)

	var book Book

	if jsonString != "false" {

		var jsonData = []byte(jsonString)
		var data Book

		var err = json.Unmarshal(jsonData, &data)
		if err != nil {

			return book, err
		}

		return data, nil

	} else {

		fmt.Println("no cache")
		err := r.db.Where("id = ?", ID).Preload("Author").Preload("Publisher").Find(&book).Error

		if err != nil {

			return book, err
		}

		jsonData, _ := json.Marshal(book)

		redisClient.SetData(r.rds, key, string(jsonData))

		return book, nil
	}

}

func (r *repository) Update(book Book) (Book, error) {

	err := r.db.Model(&book).Updates(book).Error

	if err != nil {

		return book, err
	}

	key := "book" + strconv.Itoa(book.ID)

	jsonData, _ := json.Marshal(book)

	redisClient.SetData(r.rds, key, string(jsonData))

	return book, nil

}

func (r *repository) Destroy(ID int) error {
	var book Book

	err := r.db.Where("id = ?", ID).Delete(&book).Error
	if err != nil {

		return err
	}

	return nil

}
