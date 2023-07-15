// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"time"
)

type Class struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	SchoolID  string    `json:"school_id"`
	OwnerID   string    `json:"owner_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
	School    *School   `json:"school"`
	Students  []*User   `json:"students"`
	Notes     []*Note   `json:"notes"`
}

type Comment struct {
	ID        string    `json:"id"`
	NoteID    string    `json:"note_id"`
	Comment   string    `json:"comment"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
}

type NewClass struct {
	Name     string `json:"name"`
	SchoolID string `json:"schoolID"`
	OwnerID  string `json:"ownerID"`
}

type NewComment struct {
	Comment string `json:"Comment"`
	NoteID  string `json:"NoteID"`
}

type NewJoinClass struct {
	ClassID string `json:"ClassID"`
	UserID  string `json:"UserID"`
}

type NewJoinSchool struct {
	SchoolID string `json:"SchoolID"`
	UserID   string `json:"UserID"`
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
	Name    string `json:"Name"`
	OwnerID string `json:"ownerID"`
}

type NewTag struct {
	Name string `json:"Name"`
}

type NewUser struct {
	Name     string `json:"Name"`
	Email    string `json:"Email"`
	Password string `json:"Password"`
}

type Note struct {
	ID          string     `json:"id"`
	ClassID     string     `json:"class_id"`
	SchoolID    string     `json:"school_id"`
	Description string     `json:"description"`
	Title       string     `json:"title"`
	UserID      string     `json:"user_id"`
	IsPublic    bool       `json:"is_public"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	DeletedAt   time.Time  `json:"deleted_at"`
	School      *School    `json:"school"`
	Tags        []*Tag     `json:"tags"`
	LikeUser    []*User    `json:"like_user"`
	Comment     []*Comment `json:"comment"`
}

type School struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	OwnerID   string    `json:"owner_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
	Owner     *User     `json:"owner"`
	Students  []*User   `json:"students"`
}

type Tag struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type User struct {
	ID        string    `json:"id"`
	ImageURL  string    `json:"image_url"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
	School    []*School `json:"school"`
	Likes     []*Note   `json:"likes"`
	Class     []*Class  `json:"class"`
	Notes     []*Note   `json:"notes"`
}
