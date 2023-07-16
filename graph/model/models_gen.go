// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"time"
)

type Class struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	SchoolID  string    `json:"schoolId"`
	OwnerID   string    `json:"ownerId"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	DeletedAt time.Time `json:"deletedAt"`
	School    *School   `json:"school"`
	Students  []*User   `json:"students"`
	Notes     []*Note   `json:"notes"`
}

type Comment struct {
	ID        string    `json:"id"`
	NoteID    string    `json:"noteId"`
	UserID    string    `json:"userId"`
	Comment   string    `json:"comment"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	DeletedAt time.Time `json:"deletedAt"`
}

type GetNoteProps struct {
	NoteID   *string `json:"noteID,omitempty"`
	SchoolID *string `json:"schoolID,omitempty"`
	UserID   *string `json:"userID,omitempty"`
	ClassID  *string `json:"classID,omitempty"`
	IsPublic *bool   `json:"isPublic,omitempty"`
}

type GetUserProps struct {
	Email  *string `json:"email,omitempty"`
	UserID *string `json:"userID,omitempty"`
	Name   *string `json:"Name,omitempty"`
}

type NewClass struct {
	Name     string `json:"name"`
	SchoolID string `json:"schoolID"`
	OwnerID  string `json:"ownerID"`
}

type NewComment struct {
	UserID  string `json:"userID"`
	NoteID  string `json:"noteID"`
	Comment string `json:"comment"`
}

type NewJoinClass struct {
	ClassID string `json:"classID"`
	UserID  string `json:"userID"`
}

type NewJoinSchool struct {
	SchoolID string `json:"schoolID"`
	UserID   string `json:"userID"`
}

type NewNote struct {
	ClassID     string `json:"classID"`
	SchoolID    string `json:"schoolID"`
	Description string `json:"description"`
	Title       string `json:"title"`
	UserID      string `json:"userID"`
	IsPublic    bool   `json:"isPublic"`
}

type NewSchool struct {
	Name    string `json:"name"`
	OwnerID string `json:"ownerID"`
}

type NewTag struct {
	Name string `json:"name"`
}

type NewUser struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Note struct {
	ID          string     `json:"id"`
	ClassID     string     `json:"classId"`
	SchoolID    string     `json:"schoolId"`
	Description string     `json:"description"`
	Title       string     `json:"title"`
	UserID      string     `json:"userId"`
	IsPublic    bool       `json:"isPublic"`
	CreatedAt   time.Time  `json:"createdAt"`
	UpdatedAt   time.Time  `json:"updatedAt"`
	DeletedAt   time.Time  `json:"deletedAt"`
	School      *School    `json:"school"`
	Tags        []*Tag     `json:"tags"`
	LikeUser    []*User    `json:"likeUser"`
	Comment     []*Comment `json:"comment"`
}

type School struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	OwnerID   string    `json:"ownerId"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	DeletedAt time.Time `json:"deletedAt"`
	Owner     *User     `json:"owner"`
	Students  []*User   `json:"students"`
}

type Tag struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type User struct {
	ID        string    `json:"id"`
	ImageURL  string    `json:"imageUrl"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	DeletedAt time.Time `json:"deletedAt"`
	School    []*School `json:"school"`
	Likes     []*Note   `json:"likes"`
	Class     []*Class  `json:"class"`
	Notes     []*Note   `json:"notes"`
}
