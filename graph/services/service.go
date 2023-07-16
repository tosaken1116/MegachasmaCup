package services

import (
	"context"
	"megachasma/graph/model"

	"gorm.io/gorm"
)

type UserService interface {
	CreateUser(ctx context.Context, Email string, Name string, Password string) (*model.User, error)
	GetUsersNote(ctx context.Context, userID string) ([]*model.Note, error)
	GetUsersClass(ctx context.Context, userID string) ([]*model.Class, error)
	GetUsersLike(ctx context.Context, userID string) ([]*model.Note, error)
	GetUsersSchool(ctx context.Context, userID string) ([]*model.School, error)
}

type NoteService interface {
	CreateNote(ctx context.Context, ClassID string, SchoolID string, Description string, Title string, UserID string, IsPublic bool) (*model.Note, error)
	GetNoteTags(ctx context.Context, NoteID string) ([]*model.Tag, error)
	GetLikeUserOfNote(ctx context.Context, NoteID string) ([]*model.User, error)
}

type SchoolService interface {
	GetSchoolByID(ctx context.Context, id string) (*model.School, error)
	CreateSchool(ctx context.Context, Name string, OwnerID string) (*model.School, error)
	GetSchoolBySearchWord(searchWord string) ([]*model.School, error)
}

type ClassService interface {
	CreateClass(ctx context.Context, Name string, SchoolID string, OwnerID string) (*model.Class, error)
	GetClasses(input model.GetClassesProps) ([]*model.Class, error)
}

type TagService interface {
	CreateTag(ctx context.Context, Name string) (*model.Tag, error)
	GetTags(ctx context.Context, searchWord string) ([]*model.Tag, error)
}

type CommentService interface {
	CreateComment(input model.NewComment) (*model.Comment, error)
}

type Services interface {
	UserService
	NoteService
	SchoolService
	ClassService
	TagService
	CommentService
}

type services struct {
	*userService
	*noteService
	*classService
	*schoolService
	*tagService
	*commentService
}

func New(db *gorm.DB) Services {
	return &services{
		userService:    &userService{db: db},
		noteService:    &noteService{db: db},
		schoolService:  &schoolService{db: db},
		classService:   &classService{db: db},
		tagService:     &tagService{db: db},
		commentService: &commentService{db: db},
	}
}
