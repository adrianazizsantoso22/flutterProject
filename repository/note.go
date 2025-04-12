package repository

import (
	"fmt"
	"context"
	"go-notes-taker/entity"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type NoteRepository interface {
	CreateNote(ctx context.Context, note entity.Note) (entity.Note, error)
	GetAllNotes(ctx context.Context, userID string) ([]entity.Note, error)
	GetMyNotes(ctx context.Context, userID string) ([]entity.Note, error)
	FindNoteByID(ctx context.Context, noteID string) (entity.Note, error)
	DeleteNote(ctx context.Context, noteID string) (error)
	UpdateNote(ctx context.Context, note entity.Note) (error)
}

type noteConnection struct {
	connection *gorm.DB
}

func NewNoteRepository(db *gorm.DB) NoteRepository {
	return &noteConnection{
		connection: db,
	}
}

func(db *noteConnection) CreateNote(ctx context.Context, note entity.Note) (entity.Note, error) {
	note.ID = uuid.New()
	uc := db.connection.Create(&note)
	if uc.Error != nil {
		return entity.Note{}, uc.Error
	}
	return note, nil
}

func(db *noteConnection) GetAllNotes(ctx context.Context, userID string) ([]entity.Note, error) {
	var listNotes []entity.Note
	// tx := db.connection.Preload("User").Where("is_public = ? OR creator_id = ?", true, userID).Find(&listNotes)
	tx := db.connection.Preload("User", func(db *gorm.DB) *gorm.DB {
		return db.Select("id, name")
	}).Where("is_public = ? OR creator_id = ?", true, userID).Find(&listNotes)
	
	if tx.Error != nil {
		return nil, tx.Error
	}
	return listNotes, nil
}

func(db *noteConnection) GetMyNotes(ctx context.Context, userID string) ([]entity.Note, error) {
	var listNotes []entity.Note
	tx := db.connection.Preload("User", func(db *gorm.DB) *gorm.DB {
		return db.Select("id, name")
	}).Where("creator_id = ?", userID).Find(&listNotes)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return listNotes, nil
}

func(db *noteConnection) FindNoteByID(ctx context.Context, noteID string) (entity.Note, error) {
	var note entity.Note
	ux := db.connection.Preload("User", func(db *gorm.DB) *gorm.DB {
		return db.Select("id, name")
	}).Where("id = ?", noteID).Take(&note)
	if ux.Error != nil {
        return note, fmt.Errorf("could not find note with ID %s: %w", noteID, ux.Error)
	}
	return note, nil
}

func(db *noteConnection) DeleteNote(ctx context.Context, noteID string) (error) {
    uc := db.connection.Where("id = ?", noteID).Delete(&entity.Note{})
	if uc.Error != nil {
		return uc.Error
	}
	return nil
}

func(db *noteConnection) UpdateNote(ctx context.Context, note entity.Note) (error) {
	uc := db.connection.Updates(&note)
	if uc.Error != nil {
		return uc.Error
	}
	return nil
}