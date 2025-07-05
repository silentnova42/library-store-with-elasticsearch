package model

type UserProfile struct {
	Email     string `json:"email"`
	Username  string `json:"username"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Password  string `json:"password"`
}

type LogUpUserProfile struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
