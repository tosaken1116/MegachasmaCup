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

func (cs *commentService) UpdateComment(ctx context.Context, id string, input model.UpdateCommentProps) (*model.Comment, error) {
	userID, isGet := auth.GetUserID(ctx)
	if !isGet {
		return nil, errors.New("cant get userId")
	}
	comment := new(dbModel.Comment)
	if err := cs.db.Where("id = ?", id).Find(&comment).Error; err != nil {
		return nil, err
	}
	if comment.UserID.String() != userID {
		return nil, errors.New("this comment is not yours")
	}
	if input.Comment != nil {
		comment.Comment = *input.Comment
	}
	if input.DeletedAt != nil {
		comment.DeletedAt = *input.DeletedAt
	}
	if err := cs.db.Save(&comment).Error; err != nil {
		return nil, err
	}
	createdComment := convertComment(*comment)
	return createdComment, nil
}
