package services

import (
	"megachasma/graph/model"
	dbModel "megachasma/graph/model/db"

	"gorm.io/gorm"
)

type noteService struct {
	db *gorm.DB
}

func convertNote(note dbModel.Note) *model.Note {
	return &model.Note{
		ID:          note.ID.String(),
		ClassID:     note.ClassID.String(),
		SchoolID:    note.SchoolID.String(),
		Description: note.Description,
		IsPublic:    note.IsPublic,
		UserID:      note.UserID.String(),
		CreatedAt:   note.CreatedAt,
		DeletedAt:   note.DeletedAt,
		UpdatedAt:   note.UpdatedAt,
	}
}
