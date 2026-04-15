package model

type User struct {
	UserId   string `json:"userId"`
	Email    string `json:"email"`
	Name     string `json:"name"`
	Password string `json:"password"`
}
