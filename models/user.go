package models

import "time"

type User struct {
	Id        int       `json:"id"`
	Username  string    `json:"username"`
	CreatedAt time.Time `json:"createdAt"`
	FullName  string    `json:"fullName"`
}
