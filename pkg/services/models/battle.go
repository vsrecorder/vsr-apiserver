package models

import "time"

type Battle struct {
	ID                  string    `json:"id"`
	CreatedAt           time.Time `json:"created_at"`
	UpdatedAt           time.Time `json:"updated_at"`
	GameId              string    `json:"game_id"`
	UserId              string    `json:"user_id"`
	GoFirst             bool      `json:"go_first"`
	VictoryFlg          bool      `json:"victory_flg"`
	YourPrizeCards      uint      `json:"your_prize_cards"`
	OpponentsPrizeCards uint      `json:"opponents_prize_cards"`
	Memo                string    `json:"memo"`
}
