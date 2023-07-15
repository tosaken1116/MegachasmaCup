package services

import (
	"context"
	"megachasma/graph/model"
	dbModel "megachasma/graph/model/db"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type tagService struct {
	db *gorm.DB
}

func (ts *tagService) GetTags(ctx context.Context, searchWord string) ([]*model.Tag, error) {
	tags := []*dbModel.Tag{}
	if err := ts.db.Where("name LIKE ?", "%"+searchWord+"%").Find(&tags).Error; err != nil {
		return nil, err
	}
	convertedTag := make([]*model.Tag, len(tags))
	for i, key := range tags {
		convertedTag[i] = convertTag(*key)
	}
	return convertedTag, nil
}

func convertTag(tags dbModel.Tag) *model.Tag {
	return &model.Tag{
		ID:   tags.ID.String(),
		Name: tags.Name,
	}
}

func (ts *tagService) CreateTag(ctx context.Context, Name string) (*model.Tag, error) {
	tag := dbModel.Tag{
		ID:   uuid.New(),
		Name: Name,
	}
	if err := ts.db.Create(&tag).Error; err != nil {
		return nil, err
	}
	return convertTag(tag), nil
}
