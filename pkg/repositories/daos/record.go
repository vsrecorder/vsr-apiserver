package daos

import (
	"time"

	"gorm.io/gorm"
)

type Record struct {
	ID              string `gorm:"primaryKey"`
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       gorm.DeletedAt `gorm:"index"`
	OfficialEventId uint
	UserId          string
	DeckId          string
}
