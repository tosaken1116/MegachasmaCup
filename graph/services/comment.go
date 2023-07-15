package services

import (
	"megachasma/graph/model"
	dbModel "megachasma/graph/model/db"

	"gorm.io/gorm"
)

type commentService struct {
	db *gorm.DB
}

func convertComment(comment dbModel.Comment) *model.Comment {
	return &model.Comment{
		ID:        comment.ID.String(),
		NoteID:    comment.NoteID.String(),
		Comment:   comment.Comment,
		CreatedAt: comment.CreatedAt,
		DeletedAt: comment.DeletedAt,
		UpdatedAt: comment.UpdatedAt,
	}
}
