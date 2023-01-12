package book

import (
	"errors"
)

type Service interface {
	RegisterBook(input BookDataInput) (Book, error)
	GetBookByID(input BookParamInput) (Book, error)
	UpdateBookData(inputID BookParamInput, inputData BookDataInput) (Book, error)
	DestroyBook(input BookParamInput) error
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) RegisterBook(input BookDataInput) (Book, error) {

	book := Book{}
	book.Name = input.Name
	book.AuthorID = input.AuthorID
	book.PublisherID = input.PublisherID

	newBook, err := s.repository.Create(book)
	if err != nil {
		return newBook, err
	}
	return newBook, nil
}

func (s *service) GetBookByID(input BookParamInput) (Book, error) {

	book, err := s.repository.FindById(input.ID)

	if err != nil {
		return book, err
	}

	if book.ID == 0 {

		return book, errors.New("Book not found with id")

	}
	return book, nil
}

func (s *service) UpdateBookData(inputID BookParamInput, inputData BookDataInput) (Book, error) {

	book, err := s.repository.FindById(inputID.ID)
	if err != nil {

		return book, err

	}
	book.Name = inputData.Name
	book.AuthorID = inputData.AuthorID
	book.PublisherID = inputData.PublisherID

	updateBook, err := s.repository.Update(book)

	if err != nil {
		return updateBook, err
	}
	return updateBook, nil

}

func (s *service) DestroyBook(input BookParamInput) error {

	err := s.repository.Destroy(input.ID)

	if err != nil {

		return err
	}
	return nil
}
