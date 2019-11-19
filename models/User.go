package models

import "time"

type Users struct {
	Users []User `json:"users"`
}

type User struct {
	ID        int       `json:"id,omitempty" gorm:"primary_key"`
	Name      string    `json:"name"`
	Age       int       `json:"age"`
	Posts     []Post    `json:"posts,omitempty"`
	Comments  []Comment `json:"comments,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}
