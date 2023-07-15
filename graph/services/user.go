package services

import "gorm.io/gorm"

type userService struct {
	db *gorm.DB
}
