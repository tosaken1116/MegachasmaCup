package services

import (
	"context"
	"megachasma/graph/model"

	"gorm.io/gorm"
)

type UserService interface {
	CreateUser(ctx context.Context, Email string, Name string, Password string) (*model.User, error)
}

type NoteService interface {
	CreateNote(ctx context.Context, ClassID string, SchoolID string, Description string, Title string, UserID string, IsPublic bool) (*model.Note, error)
}

type SchoolService interface {
	GetSchoolByID(ctx context.Context, id string) (*model.School, error)
}

type ClassService interface {
	CreateClass(ctx context.Context, Name string, SchoolID string, OwnerID string) (*model.Class, error)
}

type TagService interface {
	CreateTag(ctx context.Context, Name string) (*model.Tag, error)
	GetTags(ctx context.Context, searchWord string) ([]*model.Tag, error)
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
