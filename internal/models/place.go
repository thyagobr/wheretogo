package models

import "time"

type Place struct {
	ID			uint   `json:"id"`
	Name    string `json:"name"`
	Address string `json:"address"`
	Country string `json:"country"`
	City    string `json:"city"`
	Tags []Tag `json:"tags" gorm:"polymorphic:Taggable;polymorphicValue:Place"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
