package models

import "time"

type Comments struct {
	Comments []Comment `json:"comments"`
}

type Comment struct {
	ID        int       `json:"id,omitempty" gorm:"primary_key"`
	Body      string    `json:"body"`
	UserID    int       `json:"userId,omitempty"`
	User      *User     `json:"user,omitempty"`
	PostID    int       `json:"postId,omitempty"`
	Post      *Post     `json:"post,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}
