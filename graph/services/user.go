package services

import (
	"megachasma/graph/model"
	dbModel "megachasma/graph/model/db"

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
