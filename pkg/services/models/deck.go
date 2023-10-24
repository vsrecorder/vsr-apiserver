package models

import "time"

type Deck struct {
	ID             string    `json:"id"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
	UserId         string    `json:"user_id"`
	Name           string    `json:"name"`
	Code           string    `json:"code"`
	PrivateCodeFlg bool      `json:"private_code_flg"`
}
