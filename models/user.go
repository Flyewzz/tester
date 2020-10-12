package models

type User struct {
	ID       int    `json:"id"`
	Email    string `json:"email"`
	Nickname string `json:"nickname"`
	Name     string `json:"name"`
	Password string `json:"-"`
}
