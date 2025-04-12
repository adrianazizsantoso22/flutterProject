package dto

import (
	"github.com/google/uuid"
	// "go-notes-taker/entity"
)

type NoteCreateDTO struct {
	ID        		uuid.UUID   		`gorm:"primary_key;not_null;type:char(36)" json:"id" form:"id"`
	Judul 			string 				`json:"judul" form:"judul" binding:"required"`
	Konten 			string 				`json:"konten" form:"konten" binding:"required"`
	IsPublic	 	*bool  				`json:"isPublic" form:"isPublic" binding:"required"`

	CreatorID		uuid.UUID 			`gorm:"foreignKey;type:char(36);not_null" json:"CreatorID"`

	Timestamp
}

type NoteUpdateDTO struct {
	ID        		uuid.UUID   		`gorm:"primary_key" json:"id" form:"id"`
	Judul 			string 				`json:"judul" form:"judul" binding:"required"`
	Konten 			string 				`json:"konten" form:"konten" binding:"required"`
	IsPublic	 	*bool  				`json:"isPublic" form:"isPublic" binding:"required"`

	Timestamp
}