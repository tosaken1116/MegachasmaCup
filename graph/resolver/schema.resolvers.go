package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.34

import (
	"context"
	"fmt"
	"megachasma/graph/model"
	"megachasma/internal"
)

// CreateUser is the resolver for the createUser field.
func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (*model.User, error) {
	return r.Srv.CreateUser(ctx, input.Email, input.Password, input.Name)
}

// UpdateUser is the resolver for the updateUser field.
func (r *mutationResolver) UpdateUser(ctx context.Context, input *model.NewUser) (*model.User, error) {
	panic(fmt.Errorf("not implemented: UpdateUser - updateUser"))
}

// CreateNote is the resolver for the createNote field.
func (r *mutationResolver) CreateNote(ctx context.Context, input model.NewNote) (*model.Note, error) {
	return r.Srv.CreateNote(ctx, input.ClassID, input.SchoolID, input.Description, input.Title, input.UserID, input.IsPublic)
}

// UpdateNote is the resolver for the updateNote field.
func (r *mutationResolver) UpdateNote(ctx context.Context, input *model.NewNote) (*model.Note, error) {
	panic(fmt.Errorf("not implemented: UpdateNote - updateNote"))
}

// CreateClass is the resolver for the createClass field.
func (r *mutationResolver) CreateClass(ctx context.Context, input model.NewClass) (*model.Class, error) {
	return r.Srv.CreateClass(ctx, input.Name, input.SchoolID, input.OwnerID)
}

// UpdateClass is the resolver for the updateClass field.
func (r *mutationResolver) UpdateClass(ctx context.Context, input *model.NewClass) (*model.Class, error) {
	panic(fmt.Errorf("not implemented: UpdateClass - updateClass"))
}

// CreateSchool is the resolver for the createSchool field.
func (r *mutationResolver) CreateSchool(ctx context.Context, input model.NewSchool) (*model.School, error) {
	panic(fmt.Errorf("not implemented: CreateSchool - createSchool"))
}

// UpdateSchool is the resolver for the updateSchool field.
func (r *mutationResolver) UpdateSchool(ctx context.Context, input *model.NewSchool) (*model.School, error) {
	panic(fmt.Errorf("not implemented: UpdateSchool - updateSchool"))
}

// CreateComment is the resolver for the createComment field.
func (r *mutationResolver) CreateComment(ctx context.Context, input model.NewComment) (*model.Comment, error) {
	panic(fmt.Errorf("not implemented: CreateComment - createComment"))
}

// UpdateComment is the resolver for the updateComment field.
func (r *mutationResolver) UpdateComment(ctx context.Context, input *model.NewComment) (*model.Comment, error) {
	panic(fmt.Errorf("not implemented: UpdateComment - updateComment"))
}

// CreateTag is the resolver for the createTag field.
func (r *mutationResolver) CreateTag(ctx context.Context, input model.NewTag) (*model.Tag, error) {
	return r.Srv.CreateTag(ctx, input.Name)
}

// JoinClass is the resolver for the joinClass field.
func (r *mutationResolver) JoinClass(ctx context.Context, input model.NewJoinClass) (*model.Class, error) {
	panic(fmt.Errorf("not implemented: JoinClass - joinClass"))
}

// JoinSchool is the resolver for the joinSchool field.
func (r *mutationResolver) JoinSchool(ctx context.Context, input model.NewJoinSchool) (*model.School, error) {
	panic(fmt.Errorf("not implemented: JoinSchool - joinSchool"))
}

// School is the resolver for the school field.
func (r *noteResolver) School(ctx context.Context, obj *model.Note) (*model.School, error) {
	panic(fmt.Errorf("not implemented: School - school"))
}

// Tags is the resolver for the tags field.
func (r *noteResolver) Tags(ctx context.Context, obj *model.Note) ([]*model.Tag, error) {
	panic(fmt.Errorf("not implemented: Tags - tags"))
}

// LikeUser is the resolver for the like_user field.
func (r *noteResolver) LikeUser(ctx context.Context, obj *model.Note) ([]*model.User, error) {
	panic(fmt.Errorf("not implemented: LikeUser - like_user"))
}

// GetNotes is the resolver for the getNotes field.
func (r *queryResolver) GetNotes(ctx context.Context) ([]*model.Note, error) {
	panic(fmt.Errorf("not implemented: GetNotes - getNotes"))
}

// GetSchools is the resolver for the getSchools field.
func (r *queryResolver) GetSchools(ctx context.Context) ([]*model.School, error) {
	panic(fmt.Errorf("not implemented: GetSchools - getSchools"))
}

// GetClasses is the resolver for the getClasses field.
func (r *queryResolver) GetClasses(ctx context.Context) ([]*model.Class, error) {
	panic(fmt.Errorf("not implemented: GetClasses - getClasses"))
}

// GetTags is the resolver for the getTags field.
func (r *queryResolver) GetTags(ctx context.Context, searchWord string) ([]*model.Tag, error) {
	return r.Srv.GetTags(ctx, searchWord)
}

// GetMyNotes is the resolver for the getMyNotes field.
func (r *queryResolver) GetMyNotes(ctx context.Context) (*model.Note, error) {
	panic(fmt.Errorf("not implemented: GetMyNotes - getMyNotes"))
}

// GetUser is the resolver for the getUser field.
func (r *queryResolver) GetUser(ctx context.Context) (*model.User, error) {
	panic(fmt.Errorf("not implemented: GetUser - getUser"))
}

// School is the resolver for the school field.
func (r *userResolver) School(ctx context.Context, obj *model.User) ([]*model.School, error) {
	panic(fmt.Errorf("not implemented: School - school"))
}

// Likes is the resolver for the likes field.
func (r *userResolver) Likes(ctx context.Context, obj *model.User) ([]*model.Note, error) {
	panic(fmt.Errorf("not implemented: Likes - likes"))
}

// Class is the resolver for the class field.
func (r *userResolver) Class(ctx context.Context, obj *model.User) ([]*model.Class, error) {
	panic(fmt.Errorf("not implemented: Class - class"))
}

// Notes is the resolver for the notes field.
func (r *userResolver) Notes(ctx context.Context, obj *model.User) ([]*model.Note, error) {
	panic(fmt.Errorf("not implemented: Notes - notes"))
}

// Mutation returns internal.MutationResolver implementation.
func (r *Resolver) Mutation() internal.MutationResolver { return &mutationResolver{r} }

// Note returns internal.NoteResolver implementation.
func (r *Resolver) Note() internal.NoteResolver { return &noteResolver{r} }

// Query returns internal.QueryResolver implementation.
func (r *Resolver) Query() internal.QueryResolver { return &queryResolver{r} }

// User returns internal.UserResolver implementation.
func (r *Resolver) User() internal.UserResolver { return &userResolver{r} }

type mutationResolver struct{ *Resolver }
type noteResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type userResolver struct{ *Resolver }
