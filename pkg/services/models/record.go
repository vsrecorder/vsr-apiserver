package models

import "time"

type Record struct {
	ID              string    `json:"id"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
	OfficialEventId uint      `json:"official_event_id"`
	UserId          string    `json:"user_id"`
	DeckId          string    `json:"deck_id"`
}
