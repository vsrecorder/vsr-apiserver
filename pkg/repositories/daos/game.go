package daos

import (
	"time"

	"gorm.io/gorm"
)

type Game struct {
	ID                 string `gorm:"primaryKey"`
	CreatedAt          time.Time
	UpdatedAt          time.Time
	DeletedAt          gorm.DeletedAt `gorm:"index"`
	RecordId           string
	UserId             string
	OpponentsUserId    string
	BO3Flg             bool
	QualifyingRoundFlg bool
	FinalTournamentFlg bool
	VictoryFlg         bool
	OpponentsDeckInfo  string
	Memo               string
}
