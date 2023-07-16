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
		UpdatedAt: class.UpdatedAt,
		School:    &model.School{ID: class.SchoolID.String()},
		Students:  students,
		Notes:     notes,
	}
}
func (cs *classService) CreateClass(ctx context.Context, Name string, SchoolID string) (*model.Class, error) {
	userID, isGet := auth.GetUserID(ctx)
	if !isGet {
		return nil, errors.New("cant get userId")
	}
	pSchoolID, err := uuid.Parse(SchoolID)
	if err != nil {
		return nil, err
	}

	if !IsUserSchoolExist(cs.db, userID, SchoolID) {
		return nil, errors.New("you are not joined to school")
	}

	pOwnerID, err := uuid.Parse(userID)
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
	if err := cs.db.Exec("INSERT INTO class_user (user_id,class_id) VALUES (@user_id,@class_id)", sql.Named("user_id", userID), sql.Named("class_id", &newClass.ID)).Error; err != nil {
		return nil, errors.New("owner cant join class")
	}
	return convertCreateClass(newClass), nil
}

func (cs *classService) UpdateClass(ctx context.Context, id string, input model.UpdateClassProps) (*model.Class, error) {
	userID, isGet := auth.GetUserID(ctx)
	if !isGet {
		return nil, errors.New("cant get userId")
	}
	class := new(dbModel.Class)
	if err := cs.db.Where("id = ?", id).Find(&class).Error; err != nil {
		return nil, err
	}
	if userID != class.OwnerID.String() {
		return nil, errors.New("this class owner is not you")
	}

	if input.Name != nil {
		class.Name = *input.Name
	}
	if input.OwnerID != nil {
		pOwnerID, err := uuid.Parse(*input.OwnerID)
		if err != nil {
			return nil, err
		}
		class.OwnerID = pOwnerID
	}

	if err := cs.db.Save(&class).Error; err != nil {
		return nil, err
	}

	updatedClass := convertCreateClass(*class)
	return updatedClass, nil
}

func (cs *classService) GetClasses(input model.GetClassesProps) ([]*model.Class, error) {
	classes := make([]*dbModel.Class, 0)
	orm := cs.db.Model(&dbModel.Class{})

	if input.SchoolID != nil {
		orm.Where("school_id = ?", *input.SchoolID)
	}
	if input.UserID != nil {
		orm.Where("owner_id = ?", *input.UserID)
	}
	if input.ClassID != nil {
		orm.Where("id = ?", *input.ClassID)
	}
	if input.SearchWord != nil {

		orm.Where("name LIKE ?", "%"+*input.SearchWord+"%")
	}

	if err := orm.Find(&classes).Error; err != nil {
		return nil, err
	}

	result := make([]*model.Class, len(classes))
	for i, class := range classes {
		result[i] = convertCreateClass(*class)
	}

	return result, nil
}
func IsUserClassExist(db *gorm.DB, userID string, classID string) bool {
	var count int64
	db.Raw("SELECT COUNT(*) FROM class_user WHERE user_id = ? AND class_id = ?", userID, classID).Scan(&count)
	return count != 0
}
