package author

import "gorm.io/gorm"

type Repository interface {
	Create(author Author) (Author, error)
	FindById(ID int) (Author, error)
	Update(author Author) (Author, error)
	Destroy(ID int) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}

}

func (r *repository) Create(author Author) (Author, error) {
	err := r.db.Create(&author).Error

	if err != nil {
		return author, err
	}
	return author, nil
}

func (r *repository) FindById(ID int) (Author, error) {

	var author Author

	err := r.db.Where("id = ?", ID).Find(&author).Error

	if err != nil {

		return author, err
	}
	return author, nil
}

func (r *repository) Update(author Author) (Author, error) {

	err := r.db.Save(&author).Error

	if err != nil {

		return author, err
	}
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
