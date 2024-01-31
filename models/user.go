package models

import "time"

type User struct {
	Id        int       `json:"id" db:"id" config:"primarykey,serial"`
	Username  string    `json:"username" db:"username" config:"unique,notnull"`
	CreatedAt time.Time `json:"createdAt" db:"created_at" config:"default_current_timestamp"`
	FullName  string    `json:"fullName" db:"full_name"`
}
