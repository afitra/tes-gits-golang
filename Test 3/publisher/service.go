package publisher

import (
	"errors"
)

type Service interface {
	RegisterPublisher(input PublisherDataInput) (Publisher, error)
	GetPublisherByID(input PublisherParamInput) (Publisher, error)
	UpdatePublisherData(inputID PublisherParamInput, inputData PublisherDataInput) (Publisher, error)
	DestroyPublisher(input PublisherParamInput) error
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) RegisterPublisher(input PublisherDataInput) (Publisher, error) {

	publisher := Publisher{}
	publisher.Name = input.Name

	newPublisher, err := s.repository.Create(publisher)
	if err != nil {
		return newPublisher, err
	}
	return newPublisher, nil
}

func (s *service) GetPublisherByID(input PublisherParamInput) (Publisher, error) {

	publisher, err := s.repository.FindById(input.ID)
	if err != nil {
		return publisher, err
	}

	if publisher.ID == 0 {

		return publisher, errors.New("Author not found with id")

	}
	return publisher, nil
}

func (s *service) UpdatePublisherData(inputID PublisherParamInput, inputData PublisherDataInput) (Publisher, error) {

	publisher, err := s.repository.FindById(inputID.ID)
	if err != nil {

		return publisher, err

	}

	publisher.Name = inputData.Name

	updatePublisher, err := s.repository.Update(publisher)

	if err != nil {
		return updatePublisher, err
	}
	return updatePublisher, nil

}

func (s *service) DestroyPublisher(input PublisherParamInput) error {

	err := s.repository.Destroy(input.ID)

	if err != nil {

		return err
	}
	return nil
}
