package model

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

var users = make(map[string]string)
