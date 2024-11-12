package models

type User struct {
	ID       int    `json:"user_id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

type Response struct {
	StatusCode int
	Message    string
	Data       interface{}
}
