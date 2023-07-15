package services

import "gorm.io/gorm"

type noteService struct {
	db *gorm.DB
}
