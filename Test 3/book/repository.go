package book

import (
	"gorm.io/gorm"
)

type Repository interface {
	Create(book Book) (Book, error)
	FindById(ID int) (Book, error)
	Update(book Book) (Book, error)
	Destroy(ID int) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}

}

func (r *repository) Create(book Book) (Book, error) {

	err := r.db.Create(&book).Error

	if err != nil {
		return book, err
	}
	return book, nil
}

func (r *repository) FindById(ID int) (Book, error) {

	var book Book

	err := r.db.Where("id = ?", ID).Preload("Author").Preload("Publisher").Find(&book).Error

	if err != nil {

		return book, err
	}
	return book, nil
}

func (r *repository) Update(book Book) (Book, error) {

	err := r.db.Model(&book).Updates(book).Error

	if err != nil {

		return book, err
	}
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
