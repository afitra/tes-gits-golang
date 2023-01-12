package publisher

import (
	"encoding/json"
	"fmt"
	"gits/redisClient"
	"strconv"

	"github.com/go-redis/redis/v9"
	"gorm.io/gorm"
)

type Repository interface {
	Create(publisher Publisher) (Publisher, error)
	FindById(ID int) (Publisher, error)
	Update(publisher Publisher) (Publisher, error)
	Destroy(ID int) error
}

type repository struct {
	db  *gorm.DB
	rds *redis.Client
}

func NewRepository(db *gorm.DB, rds *redis.Client) *repository {
	return &repository{db, rds}

}

func (r *repository) Create(publisher Publisher) (Publisher, error) {
	err := r.db.Create(&publisher).Error

	if err != nil {
		return publisher, err
	}
	return publisher, nil
}

func (r *repository) FindById(ID int) (Publisher, error) {

	key := "publisher" + strconv.Itoa(ID)
	jsonString := redisClient.GetData(r.rds, key)

	var publisher Publisher

	if jsonString != "false" {
		fmt.Println("  cache")
		var jsonData = []byte(jsonString)
		var data Publisher

		var err = json.Unmarshal(jsonData, &data)
		if err != nil {

			return publisher, err
		}

		return data, nil
	} else {
		fmt.Println("no cache")
		err := r.db.Where("id = ?", ID).Find(&publisher).Error

		if err != nil {

			return publisher, err
		}
		jsonData, _ := json.Marshal(publisher)

		redisClient.SetData(r.rds, key, string(jsonData))

		return publisher, nil
	}

}

func (r *repository) Update(publisher Publisher) (Publisher, error) {

	err := r.db.Save(&publisher).Error

	if err != nil {

		return publisher, err
	}

	key := "publisher" + strconv.Itoa(publisher.ID)

	jsonData, _ := json.Marshal(publisher)

	redisClient.SetData(r.rds, key, string(jsonData))

	return publisher, nil

}

func (r *repository) Destroy(ID int) error {
	var publisher Publisher

	err := r.db.Where("id = ?", ID).Delete(&publisher).Error
	if err != nil {

		return err
	}

	return nil

}
