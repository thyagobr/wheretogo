package models

import "time"

type Event struct {
	ID          uint   `json:"id"`
	Name				string `json:"name"`
	StartsAt    time.Time `json:"startsAt"`
	EndsAt      *time.Time `json:"endsAt"`
	Description string `json:"description"`
	Public      bool   `json:"public"`
	Tags []Tag `json:"tags" gorm:"polymorphic:Taggable;polymorphicValue:Event"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
