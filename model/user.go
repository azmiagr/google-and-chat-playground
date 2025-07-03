package model

type GoogleUserInfo struct {
	ID      string `json:"id"`
	Email   string `json:"email"`
	Name    string `json:"name"`
	Picture string `json:"picture"`
}

type OAuthLoginResponse struct {
	Token string      `json:"token"`
	User  interface{} `json:"user"`
}

type OAuthCallbackRequest struct {
	Code  string `json:"code"`
	State string `json:"state"`
}

type UserParam struct {
	UserID   int    `json:"-"`
	Email    string `json:"-"`
	Password string `json:"-"`
	RoleID   int    `json:"-"`
}

type UserRegisterParam struct {
	Name            string `json:"name" binding:"required"`
	Email           string `json:"email" binding:"required,email"`
	Password        string `json:"password" binding:"required,min=8"`
	ConfirmPassword string `json:"confirm_password" binding:"required"`
}

type UserLoginParam struct {
	Email    string  `json:"email" binding:"required,email"`
	Password *string `json:"password" binding:"required"`
}

type LoginResponse struct {
	Token string `json:"token"`
}
