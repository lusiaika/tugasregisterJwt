package entity

import "time"

type User struct {
	Id        int       `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	Age       int       `json:"age"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UserResult struct {
	Age      int    `json:"age"`
	Email    string `json:"email"`
	Id       int    `json:"id"`
	Username string `json:"username"`
}

type UserLogin struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
