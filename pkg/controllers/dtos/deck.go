package dtos

type Deck struct {
	Name           string `json:"name"`
	Code           string `json:"code"`
	PrivateCodeFlg bool   `json:"private_code_flg"`
}
