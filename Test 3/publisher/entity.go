package publisher

import "time"

type Publisher struct {
	ID        int    `gorm:"primaryKey"`
	Name      string `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
