package services

import (
	"megachasma/graph/model"
	dbModel "megachasma/graph/model/db"

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
func (cs *commentService) CreateComment(input model.NewComment) (*model.Comment, error) {
	pUserID, err := uuid.Parse(input.NoteID)
	if err != nil {
		return nil, err
	}
	comment := dbModel.Comment{
		UserID:  pUserID,
		Comment: input.Comment,
	}
	if err := cs.db.Create(&comment).Error; err != nil {
		return nil, err
	}
	createdComment := convertComment(comment)
	return createdComment, nil
}
