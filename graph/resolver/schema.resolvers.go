package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.34

import (
	"context"
	"megachasma/graph/model"
	"megachasma/internal"
)

// ClassOwner is the resolver for the classOwner field.
func (r *classResolver) ClassOwner(ctx context.Context, obj *model.Class) (*model.User, error) {
	return r.Srv.GetClassOwner(ctx, obj.ID)
}

// ClassSchool is the resolver for the classSchool field.
func (r *classResolver) ClassSchool(ctx context.Context, obj *model.Class) (*model.School, error) {
	return r.Srv.GetClassSchool(ctx, obj.ID)
}

// ClassStudents is the resolver for the classStudents field.
func (r *classResolver) ClassStudents(ctx context.Context, obj *model.Class) ([]*model.User, error) {
	return r.Srv.GetClassStudents(ctx, obj.ID)
}

// ClassNotes is the resolver for the classNotes field.
func (r *classResolver) ClassNotes(ctx context.Context, obj *model.Class) ([]*model.Note, error) {
	return r.Srv.GetClassNotes(ctx, obj.ID)
}

// CreateUser is the resolver for the createUser field.
func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (*model.User, error) {
	return r.Srv.CreateUser(ctx, input.Email, input.Name, input.Password)
}

// UpdateUser is the resolver for the updateUser field.
func (r *mutationResolver) UpdateUser(ctx context.Context, id string, input *model.UpdateUserProps) (*model.User, error) {
	return r.Srv.UpdateUser(ctx, id, *input)
}

// CreateNote is the resolver for the createNote field.
func (r *mutationResolver) CreateNote(ctx context.Context, input model.NewNote) (*model.Note, error) {
	return r.Srv.CreateNote(ctx, input.ClassID, input.SchoolID, input.Description, input.Title, input.IsPublic)
}

// UpdateNote is the resolver for the updateNote field.
func (r *mutationResolver) UpdateNote(ctx context.Context, id string, input *model.UpdateNoteProps) (*model.Note, error) {
	return r.Srv.UpdateNote(ctx, id, *input)
}

// CreateClass is the resolver for the createClass field.
func (r *mutationResolver) CreateClass(ctx context.Context, input model.NewClass) (*model.Class, error) {
	return r.Srv.CreateClass(ctx, input.Name, input.SchoolID)
}

// UpdateClass is the resolver for the updateClass field.
func (r *mutationResolver) UpdateClass(ctx context.Context, id string, input *model.UpdateClassProps) (*model.Class, error) {
	return r.Srv.UpdateClass(ctx, id, *input)
}

// CreateSchool is the resolver for the createSchool field.
func (r *mutationResolver) CreateSchool(ctx context.Context, input model.NewSchool) (*model.School, error) {
	return r.Srv.CreateSchool(ctx, input.Name)
}

// UpdateSchool is the resolver for the updateSchool field.
func (r *mutationResolver) UpdateSchool(ctx context.Context, id string, input *model.UpdateSchoolProps) (*model.School, error) {
	return r.Srv.UpdateSchool(ctx, id, *input)
}

// CreateComment is the resolver for the createComment field.
func (r *mutationResolver) CreateComment(ctx context.Context, input model.NewComment) (*model.Comment, error) {
	return r.Srv.CreateComment(ctx, input)
}

// UpdateComment is the resolver for the updateComment field.
func (r *mutationResolver) UpdateComment(ctx context.Context, id string, input *model.UpdateCommentProps) (*model.Comment, error) {
	return r.Srv.UpdateComment(ctx, id, *input)
}

// CreateTag is the resolver for the createTag field.
func (r *mutationResolver) CreateTag(ctx context.Context, input model.NewTag) (*model.Tag, error) {
	return r.Srv.CreateTag(ctx, input.Name)
}

// JoinClass is the resolver for the joinClass field.
func (r *mutationResolver) JoinClass(ctx context.Context, classID string) (*model.Class, error) {
	return r.Srv.JoinClass(ctx, classID)
}

// JoinSchool is the resolver for the joinSchool field.
func (r *mutationResolver) JoinSchool(ctx context.Context, schoolID string) (*model.School, error) {
	return r.Srv.JoinSchool(ctx, schoolID)
}

// Like is the resolver for the like field.
func (r *mutationResolver) Like(ctx context.Context, noteID string) (*model.Note, error) {
	return r.Srv.LikeNote(ctx, noteID)
}

// DeleteLike is the resolver for the deleteLike field.
func (r *mutationResolver) DeleteLike(ctx context.Context, noteID string) (*model.Note, error) {
	return r.Srv.DeleteLikeNote(ctx, noteID)
}

// School is the resolver for the school field.
func (r *noteResolver) School(ctx context.Context, obj *model.Note) (*model.School, error) {
	return r.Srv.GetSchoolByID(ctx, obj.SchoolID)
}

// Tags is the resolver for the tags field.
func (r *noteResolver) Tags(ctx context.Context, obj *model.Note) ([]*model.Tag, error) {
	return r.Srv.GetNoteTags(ctx, obj.ID)
}

// LikeUser is the resolver for the like_user field.
func (r *noteResolver) LikeUser(ctx context.Context, obj *model.Note) ([]*model.User, error) {
	return r.Srv.GetLikeUserOfNote(ctx, obj.ID)
}

// Comments is the resolver for the comments field.
func (r *noteResolver) Comments(ctx context.Context, obj *model.Note) ([]*model.Comment, error) {
	return r.Srv.GetNoteComments(ctx, obj.ID)
}

// GetNotes is the resolver for the getNotes field.
func (r *queryResolver) GetNotes(ctx context.Context, input *model.GetNoteProps) ([]*model.Note, error) {
	return r.Srv.GetNotes(ctx, *input)
}

// GetSchools is the resolver for the getSchools field.
func (r *queryResolver) GetSchools(ctx context.Context, searchWord string) ([]*model.School, error) {
	return r.Srv.GetSchoolBySearchWord(searchWord)
}

// GetClasses is the resolver for the getClasses field.
func (r *queryResolver) GetClasses(ctx context.Context, input *model.GetClassesProps) ([]*model.Class, error) {
	return r.Srv.GetClasses(*input)
}

// GetTags is the resolver for the getTags field.
func (r *queryResolver) GetTags(ctx context.Context, searchWord string) ([]*model.Tag, error) {
	return r.Srv.GetTags(ctx, searchWord)
}

// GetUser is the resolver for the getUser field.
func (r *queryResolver) GetUser(ctx context.Context, input *model.GetUserProps) ([]*model.User, error) {
	return r.Srv.GetUser(ctx, *input)
}

// GetJwt is the resolver for the getJwt field.
func (r *queryResolver) GetJwt(ctx context.Context, input *model.GetJwtProps) (*model.Jwt, error) {
	return r.Srv.SignIn(input)
}

// SchoolOwner is the resolver for the schoolOwner field.
func (r *schoolResolver) SchoolOwner(ctx context.Context, obj *model.School) (*model.User, error) {
	return r.Srv.GetSchoolOwner(ctx, obj.ID)
}

// SchoolStudents is the resolver for the schoolStudents field.
func (r *schoolResolver) SchoolStudents(ctx context.Context, obj *model.School) ([]*model.User, error) {
	return r.Srv.GetSchoolStudents(ctx, obj.ID)
}

// UserSchool is the resolver for the userSchool field.
func (r *userResolver) UserSchool(ctx context.Context, obj *model.User) ([]*model.School, error) {
	return r.Srv.GetUsersSchool(ctx, obj.ID)
}

// UserLikes is the resolver for the userLikes field.
func (r *userResolver) UserLikes(ctx context.Context, obj *model.User) ([]*model.Note, error) {
	return r.Srv.GetUsersLike(ctx, obj.ID)
}

// UserClass is the resolver for the userClass field.
func (r *userResolver) UserClass(ctx context.Context, obj *model.User) ([]*model.Class, error) {
	return r.Srv.GetUsersClass(ctx, obj.ID)
}

// UserNotes is the resolver for the userNotes field.
func (r *userResolver) UserNotes(ctx context.Context, obj *model.User) ([]*model.Note, error) {
	return r.Srv.GetUsersNote(ctx, obj.ID)
}

// Class returns internal.ClassResolver implementation.
func (r *Resolver) Class() internal.ClassResolver { return &classResolver{r} }

// Mutation returns internal.MutationResolver implementation.
func (r *Resolver) Mutation() internal.MutationResolver { return &mutationResolver{r} }

// Note returns internal.NoteResolver implementation.
func (r *Resolver) Note() internal.NoteResolver { return &noteResolver{r} }

// Query returns internal.QueryResolver implementation.
func (r *Resolver) Query() internal.QueryResolver { return &queryResolver{r} }

// School returns internal.SchoolResolver implementation.
func (r *Resolver) School() internal.SchoolResolver { return &schoolResolver{r} }

// User returns internal.UserResolver implementation.
func (r *Resolver) User() internal.UserResolver { return &userResolver{r} }

type classResolver struct{ *Resolver }
type mutationResolver struct{ *Resolver }
type noteResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type schoolResolver struct{ *Resolver }
type userResolver struct{ *Resolver }
