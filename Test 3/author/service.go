package author

import (
	"errors"
)

type Service interface {
	RegisterAuthor(input AuthorDataInput) (Author, error)
	GetAuthorByID(input AuthorParamInput) (Author, error)
	UpdateAuthorData(inputID AuthorParamInput, inputData AuthorDataInput) (Author, error)
	DestroyAuthor(input AuthorParamInput) error
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) RegisterAuthor(input AuthorDataInput) (Author, error) {

	author := Author{}
	author.Name = input.Name

	newAuthor, err := s.repository.Create(author)
	if err != nil {
		return newAuthor, err
	}
	return newAuthor, nil
}

func (s *service) GetAuthorByID(input AuthorParamInput) (Author, error) {

	author, err := s.repository.FindById(input.ID)
	if err != nil {
		return author, err
	}

	if author.ID == 0 {

		return author, errors.New("Author not found with id")

	}
	return author, nil
}

func (s *service) UpdateAuthorData(inputID AuthorParamInput, inputData AuthorDataInput) (Author, error) {

	author, err := s.repository.FindById(inputID.ID)
	if err != nil {

		return author, err

	}

	author.Name = inputData.Name

	updateAuthor, err := s.repository.Update(author)

	if err != nil {
		return updateAuthor, err
	}
	return updateAuthor, nil

}

func (s *service) DestroyAuthor(input AuthorParamInput) error {

	err := s.repository.Destroy(input.ID)

	if err != nil {

		return err
	}
	return nil
}
