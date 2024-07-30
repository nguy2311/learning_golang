package model

import "time"

type User struct {
	UserId    string    `json:"-" db:"user_id, omitempty"`
	Fullname  string    `json:"full_name,omitempty" db:"user_id, omitempty"`
	Email     string    `json:"email,omitempty" db:"user_id, omitempty"`
	Password  string    `json:"password,omitempty" db:"user_id, omitempty"`
	Role      string    `json:"role,omitempty" db:"user_id, omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	Token     string    `json:"-"`
}
