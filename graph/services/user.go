package services

import (
	"context"
	"database/sql"
	"errors"
	"megachasma/graph/model"
	dbModel "megachasma/graph/model/db"
	"megachasma/middleware/auth"
	"megachasma/utils"

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

func (us *userService) UpdateUser(ctx context.Context, id string, input model.UpdateUserProps) (*model.User, error) {
	user := new(dbModel.User)
	if err := us.db.Where("id = ?", id).Find(&user).Error; err != nil {
		return nil, err
	}
	if input.Email != nil {
		user.Email = *input.Email
	}
	if input.Name != nil {
		user.Name = *input.Name
	}
	if input.ImageURL != nil {
		user.ImageUrl = *input.ImageURL
	}
	if err := us.db.Save(&user).Error; err != nil {
		return nil, err
	}
	updatedUser := convertUser(*user)
	return updatedUser, nil
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

func (us *userService) GetUsersLike(ctx context.Context, userID string) ([]*model.Note, error) {
	user := new(dbModel.User)
	if err := us.db.Where("id = ?", userID).Find(&user).Error; err != nil {
		return nil, err
	}
	convertedNote := make([]*model.Note, len(user.Likes))
	for i, key := range user.Likes {
		convertedNote[i] = convertNote(*key)
	}
	return convertedNote, nil
}

func (us *userService) GetUsersSchool(ctx context.Context, userID string) ([]*model.School, error) {
	user := new(dbModel.User)
	if err := us.db.Where("id = ?", userID).Find(&user).Error; err != nil {
		return nil, err
	}
	convertedSchool := make([]*model.School, len(user.School))
	for i, key := range user.School {
		convertedSchool[i] = convertSchool(*key)
	}
	return convertedSchool, nil
}

func (us *userService) GetUser(input model.GetUserProps) ([]*model.User, error) {
	user := new([]*dbModel.User)
	orm := us.db.Where("")
	if input.UserID != nil {
		orm.Where("id = ?", *input.UserID)
	} else {
		if input.Name != nil {
			orm.Where("name LIKE ?", "%"+*input.Name+"%")
		}
		if input.Email != nil {
			orm.Where("email LIKE ?", "%"+*input.Email+"%")
		}
	}
	if err := orm.Find(&user).Error; err != nil {
		return nil, err
	}
	convertedUser := make([]*model.User, len(*user))
	for i, key := range *user {
		convertedUser[i] = convertUser(*key)
	}
	return convertedUser, nil
}

func (us *userService) JoinClass(ctx context.Context, classID string) (*model.Class, error) {
	userID, isGet := auth.GetUserID(ctx)
	if !isGet {
		return nil, errors.New("cant get userId")
	}
	joinClass := new(dbModel.Class)
	if err := us.db.Where("id = ?", classID).Find(&joinClass).Error; err != nil {
		return nil, err
	}
	if err := us.db.Exec("INSERT INTO class_user (user_id,class_id) VALUES(@user_id,@class_id)", sql.Named("class_id", classID), sql.Named("user_id", userID)).Error; err != nil {
		return nil, err
	}
	return convertClass(*joinClass), nil
}

func (us *userService) JoinSchool(ctx context.Context, schoolID string) (*model.School, error) {
	userID, isGet := auth.GetUserID(ctx)
	if !isGet {
		return nil, errors.New("cant get userId")
	}
	joinSchool := new(dbModel.School)
	if err := us.db.Where("id = ?", schoolID).Find(&joinSchool).Error; err != nil {
		return nil, err
	}
	if err := us.db.Exec("INSERT INTO school_user (user_id,school_id) VALUES(@user_id,@school_id)", sql.Named("school_id", schoolID), sql.Named("user_id", userID)).Error; err != nil {
		return nil, err
	}
	return convertSchool(*joinSchool), nil
}

func (us *userService) SignIn(input *model.GetJwtProps) (*model.Jwt, error) {
	user := new(dbModel.User)
	if err := us.db.Where("email = ?", input.Email).Find(&user).Error; err != nil {
		return nil, errors.New("user not found")
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.HashedPassword), []byte(input.Password)); err != nil {
		return nil, errors.New("password is incorrect")
	}
	jwt, err := utils.GenerateJwt(user.ID.String())
	if err != nil {
		return nil, errors.New("failed to generate jwt")
	}
	return &model.Jwt{Token: jwt}, nil
}
