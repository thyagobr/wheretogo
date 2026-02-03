package models

import "time"

type User struct {
	ID           		uint   `gorm:"primaryKey"`
	Email        		string `gorm:"uniqueIndex;not null"`
	PasswordDigest 	string `gorm:"not null"`
	Token			 			string `json:"token"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
