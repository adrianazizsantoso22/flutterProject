package service

import (
	"context"
	"errors"
	"go-notes-taker/dto"
	"go-notes-taker/entity"
	"go-notes-taker/repository"

	"github.com/google/uuid"
	"github.com/mashingan/smapping"
)

type NoteService interface {
	CreateNote(ctx context.Context, noteDTO dto.NoteCreateDTO, userID uuid.UUID) (entity.Note, error)
	GetAllNotes(ctx context.Context, userID uuid.UUID) ([]entity.Note, error)
	GetMyNotes(ctx context.Context, userID uuid.UUID) ([]entity.Note, error)
	DeleteNote(ctx context.Context, userID uuid.UUID, noteID string) (error)
	UpdateNote(ctx context.Context, noteDTO dto.NoteUpdateDTO, userID uuid.UUID) (error)
}

type noteService struct {
	noteRepository repository.NoteRepository
}

func NewNoteService(nr repository.NoteRepository) NoteService {
	return &noteService{
		noteRepository: nr,
	}
}

func(ns *noteService) CreateNote(ctx context.Context, noteDTO dto.NoteCreateDTO, userID uuid.UUID) (entity.Note, error) {
	note := entity.Note{}
	err := smapping.FillStruct(&note, smapping.MapFields(noteDTO))
	note.CreatorID = userID
	if err != nil {
		return note, err
	}
	return ns.noteRepository.CreateNote(ctx, note)
}

func(ns *noteService) GetAllNotes(ctx context.Context, userID uuid.UUID) ([]entity.Note, error) {
	return ns.noteRepository.GetAllNotes(ctx, userID.String())
}

func(ns *noteService) GetMyNotes(ctx context.Context, userID uuid.UUID) ([]entity.Note, error) {
	return ns.noteRepository.GetMyNotes(ctx, userID.String())
}

func(ns *noteService) DeleteNote(ctx context.Context, userID uuid.UUID, noteID string) (error) {
	note, err := ns.noteRepository.FindNoteByID(ctx, noteID)
	if err != nil {
        return errors.New("Note tidak ditemukan")
    }

	if note.CreatorID != userID {
		return errors.New("Note bukan milikmu")
	}

	return ns.noteRepository.DeleteNote(ctx, noteID)
}

func(ns *noteService) UpdateNote(ctx context.Context, noteDTO dto.NoteUpdateDTO, userID uuid.UUID) (error) {
	note, err := ns.noteRepository.FindNoteByID(ctx, noteDTO.ID.String())
	if err != nil {
		return errors.New("Note tidak ditemukan")
	}

	if note.CreatorID != userID {
		return errors.New("Note bukan milikmu")
	}

	err = smapping.FillStruct(&note, smapping.MapFields(noteDTO))
	if err != nil {
		return errors.New("Gagal memperbarui data Note")
	}

	return ns.noteRepository.UpdateNote(ctx, note)

}