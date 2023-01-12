package book

import (
	"gits/author"
	"gits/publisher"
)

type BookFormatter struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	AuthorID    int    `json:"author_id"`
	PublisherID int    `json:"publisher_id"`
}
type BookDetailFormatter struct {
	ID          int                 `json:"id"`
	Name        string              `json:"name"`
	AuthorID    int                 `json:"author_id"`
	PublisherID int                 `json:"publisher_id"`
	Author      author.Author       `json:"author"`
	Publisher   publisher.Publisher `json:"publisher"`
}

func FormatBook(book Book) BookFormatter {

	formatter := BookFormatter{}
	formatter.ID = book.ID
	formatter.Name = book.Name
	formatter.AuthorID = book.AuthorID
	formatter.PublisherID = book.PublisherID

	return formatter
}
func FormatBookDetail(book Book) BookDetailFormatter {

	formatter := BookDetailFormatter{}
	formatter.ID = book.ID
	formatter.Name = book.Name
	formatter.AuthorID = book.AuthorID
	formatter.PublisherID = book.PublisherID
	formatter.Author = book.Author
	formatter.Publisher = book.Publisher

	return formatter
}
