package dtos

type Battle struct {
	GameId              string `json:"game_id"`
	GoFirst             bool   `json:"go_first"`
	VictoryFlg          bool   `json:"victory_flg"`
	YourPrizeCards      uint   `json:"your_prize_cards"`
	OpponentsPrizeCards uint   `json:"opponents_prize_cards"`
	Memo                string `json:"memo"`
}
