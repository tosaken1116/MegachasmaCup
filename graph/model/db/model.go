package model

import "time"

type Base struct {
	ID        string    `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	DeletedAt time.Time `json:"deleted_at"`
	UpdatedAt time.Time `json:"updated_at"  gorm:"autoUpdateTime"`
}

type User struct {
	Base
	ImageUrl string `json:"image_url"`
	Name     string `json:"name" gorm:"not null"`
	Email    string `json:"email" gorm:"unique; not null"`
	HashedPassword string `json:"hashed_password"`
	School []*School `json:"school" gorm:"many2many:school_user;"`
	Likes  []*Note   `json:"likes" gorm:"many2many:likes;"`
	Class  []*Class  `json:"class" gorm:"many2many:class_user;"`
	Notes []*Note `json:"notes"  gorm:"foreignKey:UserID;"`
}

type Note struct {
	Base
	ClassID     string `json:"class_id" gorm:"not null"`
	SchoolID    string `json:"school_id" gorm:"not null"`
	Description string `json:"description" gorm:"not null"`
	UserID      string `json:"user_id" gorm:"not null"`
	IsPublic    bool   `json:"is_public" gorm:"default false"`

	School   School     `json:"school"`
	Tags     []*Tag     `json:"tag" gorm:"many2many:tagging;"`
	LikeUser []*User    `json:"likes" gorm:"many2many:likes;"`
	Comment  []*Comment `json:"comments" gorm:"foreignKey:NoteID;"`
}

type School struct {
	Base
	Name    string `json:"name" gorm:"not null"`
	OwnerID string `json:"owner_id" gorm:"not null"`

	Owner    User    `json:"owner" gorm:"foreignKey:OwnerID"`
	Students []*User `json:"students" gorm:"many2many:school_user;"`
}

type Class struct {
	Base
	Name     string `json:"name" gorm:"not null"`
	SchoolID string `json:"school_id" gorm:"not null"`
	OwnerID  string `json:"owner_id"`

	School School  `json:"school"`
	Students   []*User `json:"students" gorm:"many2many:class_user;"`
	Notes []*Note `json:"notes" gorm:"foreignKey:ClassID"`
}

type Tag struct {
	ID   string `json:"id" gorm:"primaryKey; not null"`
	Name string `json:"name"`
}

type Comment struct {
	Base
	NoteID  string `json:"note_id" gorm:"not null"`
	Comment string `json:"comment" gorm:"not null"`
}
