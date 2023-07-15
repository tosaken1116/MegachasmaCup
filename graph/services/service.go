package services

import "gorm.io/gorm"

type UserService interface {
}

type NoteService interface {
}

type SchoolService interface {
}

type ClassService interface {
}

type TagService interface {
}

type Services interface {
	UserService
	NoteService
	SchoolService
	ClassService
	TagService
}

type services struct {
	*userService
	*noteService
	*classService
	*schoolService
	*tagService
}

func New(db *gorm.DB) Services {
	return &services{
		userService:   &userService{db: db},
		noteService:   &noteService{db: db},
		schoolService: &schoolService{db: db},
		classService:  &classService{db: db},
		tagService:    &tagService{db: db},
	}
}