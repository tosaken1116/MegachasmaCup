package services

import (
	"context"
	"errors"
	"megachasma/graph/model"
	dbModel "megachasma/graph/model/db"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type userService struct {
	db *gorm.DB
}

func convertUser(user dbModel.User) *model.User {
	return &model.User{
		ID:        user.ID.String(),
		ImageURL:  user.ImageUrl,
		Email:     user.Email,
		Name:      user.Name,
		CreatedAt: user.CreatedAt,
		DeletedAt: user.DeletedAt,
		UpdatedAt: user.UpdatedAt,
	}
}

func (us *userService) CreateUser(ctx context.Context, Email string, Name string, Password string) (*model.User, error) {
	if err := us.db.Where(&dbModel.User{Email: Email}).First(&dbModel.User{}).Error; err == nil {
		return nil, errors.New("email is already in use")
	}
	hashed, _ := bcrypt.GenerateFromPassword([]byte(Password), 10)
	newUser := dbModel.User{Name: Name, Email: Email, HashedPassword: string(hashed)}
	if err := us.db.Create(&newUser).Error; err != nil {
		return nil, errors.New("failed create user")
	}
	return convertUser(newUser), nil
}

func (us *userService) GetUsersNote(ctx context.Context, userID string) ([]*model.Note, error) {
	note := new([]*dbModel.Note)
	if err := us.db.Where("user_id = ?", userID).Find(&note).Error; err != nil {
		return nil, err
	}
	convertedNote := make([]*model.Note, len(*note))
	for i, key := range *note {
		convertedNote[i] = convertNote(*key)
	}
	return convertedNote, nil
}

func (us *userService) GetUsersClass(ctx context.Context, userID string) ([]*model.Class, error) {
	user := new(dbModel.User)
	if err := us.db.Where("id = ?", userID).Find(&user).Error; err != nil {
		return nil, err
	}
	convertedClass := make([]*model.Class, len(user.Class))
	for i, key := range user.Class {
		convertedClass[i] = convertClass(*key)
	}
	return convertedClass, nil
}
