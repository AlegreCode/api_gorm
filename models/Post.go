package models

import "time"

type Posts struct {
	Posts []Post `json:"posts"`
}

type Post struct {
	ID        int       `json:"id,omitempty" gorm:"primary_key"`
	Title     string    `json:"title"`
	Body      string    `json:"body"`
	UserID    int       `json:"userId,omitempty"`
	User      *User     `json:"user,omitempty"`
	Comments  []Comment `json:"comments,omitempty"`
	CreatedAt time.Time `json:"created_at,emitempty"`
	UpdatedAt time.Time `json:"updated_at,emitempty"`
}
