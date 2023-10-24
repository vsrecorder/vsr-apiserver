package daos

import (
	"time"

	"gorm.io/gorm"
)

type Deck struct {
	ID             string `gorm:"primaryKey"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      gorm.DeletedAt `gorm:"index"`
	UserId         string
	Name           string
	Code           string
	PrivateCodeFlg bool
}
