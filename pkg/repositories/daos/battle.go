package daos

import (
	"time"

	"gorm.io/gorm"
)

type Battle struct {
	ID                  string `gorm:"primaryKey"`
	CreatedAt           time.Time
	UpdatedAt           time.Time
	DeletedAt           gorm.DeletedAt `gorm:"index"`
	GameId              string
	UserId              string
	GoFirst             bool
	VictoryFlg          bool
	YourPrizeCards      uint
	OpponentsPrizeCards uint
	Memo                string
}
