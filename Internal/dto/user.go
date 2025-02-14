package dto

type UserRegisterRequest struct {
	Name     string `json:"name"`
	Nickname string `json:"nickname"`
	Bio      string `json:"bio"`

	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UserRegisterResponse struct {
	Message string `json:"message"`
}

type UserUsernameAvailRequest struct {
	Username string `json:"username" binding:"required"`
}