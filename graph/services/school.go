package services

import (
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
func GetSchoolByID(db *gorm.DB, id string) (*dbModel.School, error) {
	school := new(dbModel.School)
	if err := db.Where("id = ?", id).Find(&school).Error; err != nil {
		return nil, err
	}
	return school, nil
}

func (ss *schoolService) UpdateSchool(input model.School) (*model.School, error) {
	school, err := GetSchoolByID(ss.db, input.ID)
	if err != nil {
		return nil, err
	}
	school.Name = input.Name
	if err := ss.db.Save(&school).Error; err != nil {
		return nil, err
	}
	updatedSchool := convertSchool(*school)
	return updatedSchool, nil
}
