package publisher

import "gorm.io/gorm"

type Repository interface {
	Create(publisher Publisher) (Publisher, error)
	FindById(ID int) (Publisher, error)
	Update(publisher Publisher) (Publisher, error)
	Destroy(ID int) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}

}

func (r *repository) Create(publisher Publisher) (Publisher, error) {
	err := r.db.Create(&publisher).Error

	if err != nil {
		return publisher, err
	}
	return publisher, nil
}

func (r *repository) FindById(ID int) (Publisher, error) {

	var publisher Publisher

	err := r.db.Where("id = ?", ID).Find(&publisher).Error

	if err != nil {

		return publisher, err
	}
	return publisher, nil
}

func (r *repository) Update(publisher Publisher) (Publisher, error) {

	err := r.db.Save(&publisher).Error

	if err != nil {

		return publisher, err
	}
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
