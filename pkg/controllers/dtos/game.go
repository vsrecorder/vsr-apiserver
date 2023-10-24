package dtos

type Game struct {
	RecordId           string `json:"record_id"`
	OpponentsUserId    string `json:"opponentes_user_id"`
	BO3Flg             bool   `json:"bo3_flg"`
	QualifyingRoundFlg bool   `json:"qualifying_round_flg"`
	FinalTournamentFlg bool   `json:"final_tournament_flg"`
	VictoryFlg         bool   `json:"victory_flg"`
	OpponentsDeckInfo  string `json:"opponents_deck_info"`
	Memo               string `json:"memo"`
}
