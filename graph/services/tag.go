package services

import (
	"context"
	"megachasma/graph/model"
	dbModel "megachasma/graph/model/db"

	"gorm.io/gorm"
)

type tagService struct {
	db *gorm.DB
}

func (ts *tagService) GetTags(ctx context.Context, searchWord string) (*model.Tag, error) {
	tags := dbModel.Tag{}
	if err := ts.db.Find(&tags).Error; err != nil {
		return nil, err
	}
	return convertTag(tags), nil
}

func convertTag(tags dbModel.Tag) *model.Tag {
	return &model.Tag{
		ID:   tags.ID,
		Name: tags.Name,
	}
}
