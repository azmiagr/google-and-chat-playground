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
