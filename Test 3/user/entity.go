package user

import "time"

type User struct {
	ID           int    `gorm:"primaryKey"`
	Name         string `gorm:"not null"`
	Email        string `gorm:"not null"`
	PasswordHash string `gorm:"not null"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
