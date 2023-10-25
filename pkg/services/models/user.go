package models

type User struct {
	UID         string `json:"uid"`
	DisplayName string `json:"display_name"`
	PhotoURL    string `json:"photo_url"`
}
