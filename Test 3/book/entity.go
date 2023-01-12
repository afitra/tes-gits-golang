package book

import (
	"gits/author"
	"gits/publisher"
	"time"

	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	ID          int    `gorm:"primaryKey"`
	Name        string `gorm:"not null"`
	AuthorID    int    `gorm:"not null"`
	PublisherID int    `gorm:"not null"`
	Author      author.Author
	Publisher   publisher.Publisher
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
