package models

import "time"

type User struct {
	Id        int       `json:"id"`
	UserName  string    `json:"username"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"createdAt"`
}
