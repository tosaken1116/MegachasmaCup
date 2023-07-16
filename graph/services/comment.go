package services

import (
	"context"
	"errors"
	"megachasma/graph/model"
	dbModel "megachasma/graph/model/db"
	"megachasma/middleware/auth"

	"github.com/google/uuid"
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
func (cs *commentService) CreateComment(ctx context.Context, input model.NewComment) (*model.Comment, error) {
	userID, isGet := auth.GetUserID(ctx)
	if !isGet {
		return nil, errors.New("cant get userId")
	}
	pUserID, err := uuid.Parse(userID)
	if err != nil {
		return nil, err
	}
	pNoteID, err := uuid.Parse(input.NoteID)
	if err != nil {
		return nil, err
	}

	comment := dbModel.Comment{
		UserID:  pUserID,
		NoteID:  pNoteID,
		Comment: input.Comment,
	}
	if err := cs.db.Create(&comment).Error; err != nil {
		return nil, err
	}
	createdComment := convertComment(comment)
	return createdComment, nil
}
