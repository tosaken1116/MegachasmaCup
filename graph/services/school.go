package services

import (
	"context"
	"database/sql"
	"errors"
	"megachasma/graph/model"
	dbModel "megachasma/graph/model/db"
	"megachasma/middleware/auth"

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

func (ss *schoolService) CreateSchool(ctx context.Context, Name string) (*model.School, error) {
	userID, isGet := auth.GetUserID(ctx)
	if !isGet {
		return nil, errors.New("cant get userId")
	}
	pOwnerID, err := uuid.Parse(userID)
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
	if err := ss.db.Exec("INSERT INTO school_user (user_id,school_id) VALUES (@user_id,@school_id)", sql.Named("user_id", userID), sql.Named("school_id", &school.ID)).Error; err != nil {
		return nil, errors.New("owner cant join school")
	}
	createdSchool := convertSchool(school)
	return createdSchool, nil
}

func (ss *schoolService) UpdateSchool(ctx context.Context, id string, input model.UpdateSchoolProps) (*model.School, error) {

	school := new(dbModel.School)
	if err := ss.db.Where("id = ?", id).Find(&school).Error; err != nil {
		return nil, err
	}
	if input.OwnerID != nil {
		pOwnerID, err := uuid.Parse(*input.OwnerID)
		if err != nil {
			return nil, err
		}
		school.OwnerID = pOwnerID
	}
	if input.Name != nil {
		school.Name = *input.Name
	}

	if err := ss.db.Save(&school).Error; err != nil {
		return nil, err
	}

	updatedSchool := convertSchool(*school)
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
