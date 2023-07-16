package services

import (
	"context"
	"megachasma/graph/model"
	dbModel "megachasma/graph/model/db"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type noteService struct {
	db *gorm.DB
}

func convertNote(note dbModel.Note) *model.Note {
	return &model.Note{
		ID:          note.ID.String(),
		ClassID:     note.ClassID.String(),
		SchoolID:    note.SchoolID.String(),
		Description: note.Description,
		IsPublic:    note.IsPublic,
		UserID:      note.UserID.String(),
		CreatedAt:   note.CreatedAt,
		DeletedAt:   note.DeletedAt,
		UpdatedAt:   note.UpdatedAt,
	}
}
func convertCreateNote(note dbModel.Note) *model.Note {
	likeUser := make([]*model.User, len(note.LikeUser))
	for i, key := range note.LikeUser {
		likeUser[i] = convertUser(*key)
	}
	tags := make([]*model.Tag, len(note.Tags))
	for i, key := range note.Tags {
		tags[i] = convertTag(*key)
	}
	comment := make([]*model.Comment, len(note.Comment))
	for i, key := range note.Comment {
		comment[i] = convertComment(*key)
	}
	return &model.Note{
		ID:          note.ID.String(),
		ClassID:     note.ClassID.String(),
		SchoolID:    note.SchoolID.String(),
		Title:       note.Title,
		Description: note.Description,
		UserID:      note.UserID.String(),
		IsPublic:    note.IsPublic,
		School:      &model.School{ID: note.SchoolID.String()},
		Tags:        tags,
		LikeUser:    likeUser,
		Comment:     comment,
	}
}

func (ns *noteService) CreateNote(ctx context.Context, ClassID string, SchoolID string, Description string, Title string, UserID string, IsPublic bool) (*model.Note, error) {
	pSchoolID, err := uuid.Parse(SchoolID)
	if err != nil {
		return nil, err
	}
	pClassID, err := uuid.Parse(ClassID)
	if err != nil {
		return nil, err
	}
	pUserID, err := uuid.Parse(UserID)
	if err != nil {
		return nil, err
	}
	newNote := dbModel.Note{
		ClassID:     pClassID,
		SchoolID:    pSchoolID,
		Title:       Title,
		Description: Description,
		UserID:      pUserID,
		IsPublic:    IsPublic,
	}
	if err := ns.db.Create(&newNote).Error; err != nil {
		return nil, err
	}
	return convertCreateNote(newNote), nil
}

func (ns *noteService) GetNoteTags(ctx context.Context, NoteID string) ([]*model.Tag, error) {
	note := new(dbModel.Note)
	if err := ns.db.Where("id = ?", NoteID).Find(&note).Error; err != nil {
		return nil, err
	}
	convertedTag := make([]*model.Tag, len(note.Tags))
	for i, key := range note.Tags {
		convertedTag[i] = convertTag(*key)
	}
	return convertedTag, nil
}
func (ns *noteService) GetLikeUserOfNote(ctx context.Context, NoteID string) ([]*model.User, error) {
	note := new(dbModel.Note)
	if err := ns.db.Where("id = ?", NoteID).Find(&note).Error; err != nil {
		return nil, err
	}
	convertedUser := make([]*model.User, len(note.LikeUser))
	for i, key := range note.LikeUser {
		convertedUser[i] = convertUser(*key)
	}
	return convertedUser, nil
}

func (ns *noteService) GetNotes(input model.GetNoteProps) ([]*model.Note, error) {
	note := new([]*dbModel.Note)
	orm := ns.db.Where("")
	if input.NoteID != nil {
		orm.Where("id = ?", *input.NoteID)
	} else {
		if input.ClassID != nil {
			orm.Where("class_id = ?", *input.ClassID)
		}
		if input.SchoolID != nil {
			orm.Where("school_id = ?", *input.SchoolID)
		}
		if input.IsPublic != nil {
			orm.Where("is_public = ?", *input.IsPublic)
		}

	}
	if err := orm.Find(&note).Error; err != nil {
		return nil, err
	}
	convertedNote := make([]*model.Note, len(*note))
	for i, key := range *note {
		convertedNote[i] = convertNote(*key)
	}
	return convertedNote, nil
}
