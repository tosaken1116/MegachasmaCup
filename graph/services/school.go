package services

import (
	"context"
	"megachasma/graph/model"
	dbModel "megachasma/graph/model/db"

	"gorm.io/gorm"
)

type schoolService struct {
	db *gorm.DB
}

func convertSchool(school dbModel.School) *model.School {
	return &model.School{
		ID:        school.ID.String(),
		Name:      school.Name,
		OwnerID:   school.OwnerID.String(),
		CreatedAt: school.CreatedAt,
		DeletedAt: school.DeletedAt,
		UpdatedAt: school.UpdatedAt,
	}
}
func (ss *schoolService) GetSchoolByID(ctx context.Context, id string) (*model.School, error) {
	school := new(dbModel.School)
	if err := ss.db.Where("id = ?", id).Find(&school).Error; err != nil {
		return nil, err
	}
	return convertSchool(*school), nil
}
