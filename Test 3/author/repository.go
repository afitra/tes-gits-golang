package author

import (
	"encoding/json"
	"gits/redisClient"
	"strconv"

	"github.com/go-redis/redis/v9"
	"gorm.io/gorm"
)

type Repository interface {
	Create(author Author) (Author, error)
	FindById(ID int) (Author, error)
	Update(author Author) (Author, error)
	Destroy(ID int) error
}

type repository struct {
	db  *gorm.DB
	rds *redis.Client
}

func NewRepository(db *gorm.DB, rds *redis.Client) *repository {
	return &repository{db, rds}

}

func (r *repository) Create(author Author) (Author, error) {

	err := r.db.Create(&author).Error

	if err != nil {
		return author, err
	}
	return author, nil

}

func (r *repository) FindById(ID int) (Author, error) {

	key := "author" + strconv.Itoa(ID)
	jsonString := redisClient.GetData(r.rds, key)

	var author Author
	if jsonString != "false" {

		var jsonData = []byte(jsonString)
		var data Author

		var err = json.Unmarshal(jsonData, &data)
		if err != nil {

			return author, err
		}

		return data, nil

	} else {

		err := r.db.Where("id = ?", ID).Find(&author).Error

		if err != nil {

			return author, err
		}
		jsonData, _ := json.Marshal(author)

		redisClient.SetData(r.rds, key, string(jsonData))

		return author, nil
	}

}

func (r *repository) Update(author Author) (Author, error) {

	err := r.db.Save(&author).Error

	if err != nil {

		return author, err
	}

	key := "author" + strconv.Itoa(author.ID)

	jsonData, _ := json.Marshal(author)

	redisClient.SetData(r.rds, key, string(jsonData))

	return author, nil

}

func (r *repository) Destroy(ID int) error {
	var author Author

	err := r.db.Where("id = ?", ID).Delete(&author).Error
	if err != nil {

		return err
	}

	return nil

}
