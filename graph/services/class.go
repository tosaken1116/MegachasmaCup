package services

import (
	"context"
	"megachasma/graph/model"
	dbModel "megachasma/graph/model/db"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type classService struct {
	db *gorm.DB
}

func convertClass(class dbModel.Class) *model.Class {
	return &model.Class{
		ID:       class.ID.String(),
		Name:     class.Name,
		SchoolID: class.SchoolID.String(),
		OwnerID:  class.OwnerID.String(),
	}
}
func convertCreateClass(class dbModel.Class) *model.Class {
	students := make([]*model.User, len(class.Students))
	for i, key := range class.Students {
		students[i] = convertUser(*key)
	}
	notes := make([]*model.Note, len(class.Notes))
	for i, key := range class.Notes {
		notes[i] = convertNote(*key)
	}
	return &model.Class{
		ID:        class.ID.String(),
		Name:      class.Name,
		SchoolID:  class.SchoolID.String(),
		OwnerID:   class.OwnerID.String(),
		CreatedAt: class.CreatedAt,
		DeletedAt: class.DeletedAt,
		UpdatedAt: class.UpdatedAt,
		School:    &model.School{ID: class.SchoolID.String()},
		Students:  students,
		Notes:     notes,
	}
}
func (cs *classService) CreateClass(ctx context.Context, Name string, SchoolID string, OwnerID string) (*model.Class, error) {
	pSchoolID, err := uuid.Parse(SchoolID)
	if err != nil {
		return nil, err
	}
	pOwnerID, err := uuid.Parse(OwnerID)
	if err != nil {
		return nil, err
	}
	newClass := dbModel.Class{
		Name:     Name,
		SchoolID: pSchoolID,
		OwnerID:  pOwnerID,
	}
	if err := cs.db.Create(&newClass).Error; err != nil {
		return nil, err
	}
	return convertCreateClass(newClass), nil
}

func (cs *classService) GetClasses(ctx context.Context) ([]*model.Class, error) {
	var classes []*dbModel.Class
	if err := cs.db.Find(&classes).Error; err != nil {
		return nil, err
	}

	result := make([]*model.Class, len(classes))
	for i, class := range classes {
		result[i] = convertCreateClass(*class)
	}

	return result, nil
}
