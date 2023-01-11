package user

import "time"

type User struct {
	ID           int `gorm:"primaryKey"`
	Name         string
	Email        string
	PasswordHash string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
