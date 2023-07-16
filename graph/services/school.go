package services

import (
	"context"
	"megachasma/graph/model"
	dbModel "megachasma/graph/model/db"

	"github.com/google/uuid"
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

func (ss *schoolService) CreateSchool(ctx context.Context, Name string, OwnerID string) (*model.School, error) {
	pOwnerID, err := uuid.Parse(OwnerID)
	if err != nil {
		return nil, err
	}
	school := dbModel.School{
		Name:    Name,
		OwnerID: pOwnerID,
	}
	if err := ss.db.Create(&school).Error; err != nil {
		return nil, err
	}
	createdSchool := convertSchool(school)
	return createdSchool, nil
}

func (ss *schoolService) UpdateSchool(ctx context.Context, Name string, OwnerID string) (*model.School, error) {
	pOwnerID, err := uuid.Parse(OwnerID)
	if err != nil {
		return nil, err
	}
	school := dbModel.School{
		Name:    Name,
		OwnerID: pOwnerID,
	}
	if err := ss.db.Where("id = ?", OwnerID).Find(&school).Error; err != nil {
		return nil, err
	}

	school.Name = Name

	if err := ss.db.Save(&school).Error; err != nil {
		return nil, err
	}

	updatedSchool := convertSchool(school)
	return updatedSchool, nil
}

func (ss *schoolService) GetSchoolBySearchWord(searchWord string) ([]*model.School, error) {
	schools := []*dbModel.School{}
	if err := ss.db.Where("name LIKE ?", "%"+searchWord+"%").Find(&schools).Error; err != nil {
		return nil, err
	}
	convertedSchool := make([]*model.School, len(schools))
	for i, key := range schools {
		convertedSchool[i] = convertSchool(*key)
	}
	return convertedSchool, nil
}
