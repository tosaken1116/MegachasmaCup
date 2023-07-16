package services

import (
	"context"
	"database/sql"
	"errors"
	"megachasma/graph/model"
	dbModel "megachasma/graph/model/db"
	"megachasma/middleware/auth"

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

func (ns *noteService) CreateNote(ctx context.Context, ClassID string, SchoolID string, Description string, Title string, IsPublic bool) (*model.Note, error) {
	userID, isGet := auth.GetUserID(ctx)
	if !isGet {
		return nil, errors.New("cant get userId")
	}
	pUserID, err := uuid.Parse(userID)
	if err != nil {
		return nil, err
	}
	pSchoolID, err := uuid.Parse(SchoolID)
	if err != nil {
		return nil, err
	}
	pClassID, err := uuid.Parse(ClassID)
	if err != nil {
		return nil, err
	}
	if !IsUserSchoolExist(ns.db, userID, SchoolID) {
		return nil, errors.New("you are not joined to school")
	}
	if !IsUserClassExist(ns.db, userID, ClassID) {
		return nil, errors.New("you are not joined to class")
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
	if err := ns.db.Preload("Tags").Where("id = ?", NoteID).Find(&note).Error; err != nil {
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
	if err := ns.db.Preload("LikeUser").Where("id = ?", NoteID).Find(&note).Error; err != nil {
		return nil, err
	}
	convertedUser := make([]*model.User, len(note.LikeUser))
	for i, key := range note.LikeUser {
		convertedUser[i] = convertUser(*key)
	}
	return convertedUser, nil
}

func (ns *noteService) GetNotes(ctx context.Context, input model.GetNoteProps) ([]*model.Note, error) {
	userID, isGet := auth.GetUserID(ctx)
	if !isGet {
		return nil, errors.New("cant get userId")
	}

	note := new([]*dbModel.Note)
	orm := ns.db.Where("")
	if input.IsMy != nil {
		if *input.IsMy {
			if input.UserID != nil {
				return nil, errors.New("isMy and userId cannot use the same time")
			}
			if input.IsPublic != nil {
				orm.Where("is_public = ?", *input.IsPublic)
			}
		}
	} else if input.ClassID != nil && input.SchoolID != nil {
		if !IsUserSchoolExist(ns.db, userID, *input.SchoolID) {
			return nil, errors.New("you are not joined to school")
		}
		if !IsUserClassExist(ns.db, userID, *input.ClassID) {
			return nil, errors.New("you are not joined to class")
		}
		if input.UserID != nil {
			orm.Where("user_id = ?", *input.UserID)
		}
		orm.Where("class_id = ?", *input.ClassID)
		orm.Where("school_id = ?", *input.SchoolID)
		orm.Where("is_public = ?", true)
	} else {
		return nil, errors.New("class id and school id or isMy is required")
	}
	if input.NoteID != nil {
		orm.Where("id = ?", *input.NoteID)
	}
	if err := orm.Find(&note).Error; err != nil {
		return nil, err
	}
	existMyNote := false
	convertedNote := make([]*model.Note, len(*note))
	for i, key := range *note {
		if key.UserID.String() == userID {
			existMyNote = true
		}
		convertedNote[i] = convertNote(*key)
	}
	if !existMyNote {
		return nil, errors.New("not allowed because your note is not found")
	}
	return convertedNote, nil
}
func (ns *noteService) LikeNote(ctx context.Context, noteID string) (*model.Note, error) {
	userID, isGet := auth.GetUserID(ctx)
	if !isGet {
		return nil, errors.New("cant get userId")
	}
	likeNote := new(dbModel.Note)
	if err := ns.db.Where("id = ?", noteID).Find(&likeNote).Error; err != nil {
		return nil, err
	}
	if err := ns.db.Exec("INSERT INTO likes (user_id,note_id) VALUES(@user_id,@note_id)", sql.Named("note_id", noteID), sql.Named("user_id", userID)).Error; err != nil {
		return nil, err
	}
	return convertNote(*likeNote), nil
}
func (ns *noteService) DeleteLikeNote(ctx context.Context, noteID string) (*model.Note, error) {
	userID, isGet := auth.GetUserID(ctx)
	if !isGet {
		return nil, errors.New("cant get userId")
	}
	likeNote := new(dbModel.Note)
	if err := ns.db.Where("id = ?", noteID).Find(&likeNote).Error; err != nil {
		return nil, err
	}
	if count := ns.db.Exec("DELETE FROM likes WHERE user_id = @user_id AND note_id = @note_id", sql.Named("note_id", noteID), sql.Named("user_id", userID)).RowsAffected; count == 0 {
		return nil, errors.New("yet like note")
	}
	return convertNote(*likeNote), nil
}

func (ns *noteService) UpdateNote(ctx context.Context, id string, input model.UpdateNoteProps) (*model.Note, error) {
	userID, isGet := auth.GetUserID(ctx)
	if !isGet {
		return nil, errors.New("cant get userId")
	}
	note := new(dbModel.Note)
	if err := ns.db.Where("id = ?", id).Find(&note).Error; err != nil {
		return nil, err
	}
	if userID != note.UserID.String() {
		return nil, errors.New("this note is not yours")
	}
	if input.Description != nil {
		note.Description = *input.Description
	}
	if input.IsPublic != nil {
		note.IsPublic = *input.IsPublic
	}
	if input.Title != nil {
		note.Title = *input.Title
	}
	if input.DeletedAt != nil {
		note.DeletedAt = *input.DeletedAt
	}
	if err := ns.db.Save(&note).Error; err != nil {
		return nil, err
	}
	return convertNote(*note), nil
}
