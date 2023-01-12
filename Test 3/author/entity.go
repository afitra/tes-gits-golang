package author

import "time"

type Author struct {
	ID        int    `gorm:"primaryKey"`
	Name      string `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
