package dto

type UserRegisterForm struct {
	Name     string `json:"name"`
	Nickname string `json:"nickname"`
	Bio      string `json:"bio"`

	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}