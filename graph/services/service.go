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
	GetUser(ctx context.Context, input model.GetUserProps) ([]*model.User, error)
	JoinClass(ctx context.Context, classID string) (*model.Class, error)
	JoinSchool(ctx context.Context, schoolID string) (*model.School, error)
	SignIn(input *model.GetJwtProps) (*model.Jwt, error)
	UpdateUser(ctx context.Context, id string, input model.UpdateUserProps) (*model.User, error)
	GetSchoolOwner(ctx context.Context, SchoolID string) (*model.User, error)
	GetSchoolStudents(ctx context.Context, SchoolID string) ([]*model.User, error)
	GetClassOwner(ctx context.Context, ClassID string) (*model.User, error)
	GetClassStudents(ctx context.Context, ClassID string) ([]*model.User, error)
}

type NoteService interface {
	CreateNote(ctx context.Context, ClassID string, SchoolID string, Description string, Title string, IsPublic bool) (*model.Note, error)
	GetNoteTags(ctx context.Context, NoteID string) ([]*model.Tag, error)
	GetLikeUserOfNote(ctx context.Context, NoteID string) ([]*model.User, error)
	GetNotes(ctx context.Context, input model.GetNoteProps) ([]*model.Note, error)
	LikeNote(ctx context.Context, noteID string) (*model.Note, error)
	DeleteLikeNote(ctx context.Context, noteID string) (*model.Note, error)
	UpdateNote(ctx context.Context, id string, input model.UpdateNoteProps) (*model.Note, error)
	GetNoteComments(ctx context.Context, noteId string) ([]*model.Comment, error)
	GetClassNotes(ctx context.Context, classID string) ([]*model.Note, error)
}

type SchoolService interface {
	GetSchoolByID(ctx context.Context, id string) (*model.School, error)
	CreateSchool(ctx context.Context, Name string) (*model.School, error)
	GetSchoolBySearchWord(searchWord string) ([]*model.School, error)
	UpdateSchool(ctx context.Context, id string, input model.UpdateSchoolProps) (*model.School, error)
	GetUsersSchool(ctx context.Context, userID string) ([]*model.School, error)
	GetClassSchool(ctx context.Context, classID string) (*model.School, error)
}

type ClassService interface {
	CreateClass(ctx context.Context, Name string, SchoolID string) (*model.Class, error)
	GetClasses(input model.GetClassesProps) ([]*model.Class, error)
	UpdateClass(ctx context.Context, id string, input model.UpdateClassProps) (*model.Class, error)
	GetUserClass(ctx context.Context, userID string) ([]*model.Class, error)
}

type TagService interface {
	CreateTag(ctx context.Context, Name string) (*model.Tag, error)
	GetTags(ctx context.Context, searchWord string) ([]*model.Tag, error)
}

type CommentService interface {
	CreateComment(ctx context.Context, input model.NewComment) (*model.Comment, error)
	UpdateComment(ctx context.Context, id string, input model.UpdateCommentProps) (*model.Comment, error)
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
