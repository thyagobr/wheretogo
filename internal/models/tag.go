package models

import "time"

type Tag struct {
	ID        		uint      `json:"id"`
	Text      		string    `json:"text"`
	TaggableID   	uint      `json:"taggable_id"`
	TaggableType 	string    `json:"taggable_type"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
