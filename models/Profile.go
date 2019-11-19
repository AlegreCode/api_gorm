package models

import "time"

type Profiles struct {
	Profiles []Profile `json:"profiles"`
}

type Profile struct {
	ID        int       `json:"id,omitempty" gorm:"primary_key"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	UserID    int       `json:"userId,omitempty"`
	User      *User     `json:"user,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}
