package models

import "time"

type Event struct {
	ID          uint       `json:"id"`
	Name        string     `json:"name"`
	StartsAt    time.Time  `json:"startsAt" gorm:"column:start_at"`
	EndsAt      *time.Time `json:"endsAt" gorm:"column:end_at"`
	Description string     `json:"description"`
	Public      bool       `json:"public"`
	Tags        []Tag      `json:"tags" gorm:"polymorphic:Taggable;polymorphicValue:Event"`
	PlaceID     uint       `json:"place_id"`                        // foreign key column
	Place       Place      `json:"place" gorm:"foreignKey:PlaceID"` // association
	UserID      uint       `json:"user_id"`                         // foreign key column
	User        User       `json:"user" gorm:"foreignKey:UserID"`   // association
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
}
