package entity

import (
	"github.com/google/uuid"
)

type Note struct {
	ID        		uuid.UUID   `gorm:"primary_key;not_null" json:"id"`
	Judul 			string 		`json:"judul"`
	Konten 			string 		`json:"konten"`
	IsPublic		bool		`json:"isPublic"`

	CreatorID		uuid.UUID	`gorm:"foreignKey:CreatorID;references:ID;not_null" json:"CreatorID"`
	User 			User 		`gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:CreatorID;references:ID" json:"user,omitempty"`
	
	Timestamp
}