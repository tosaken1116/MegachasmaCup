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
func (cs *commentService) CreateComment(input model.NewComment) (*model.Comment, error) {
	comment := dbModel.Comment{
		UserID:  input.UserID,
		NoteID:  input.NoteID,
		Comment: input.Comment,
	}
	if err := cs.db.Create(&comment).Error; err != nil {
		return nil, err
	}
	createdComment := convertComment(comment)
	return createdComment, nil
}
