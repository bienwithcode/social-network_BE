package model

type AuthEmailPassword struct {
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

type Token struct {
	Token string `json:"token"`
	// ExpiredIn in seconds
	ExpiredIn int64 `json:"expire_in"`
}

type TokenResponse struct {
	AccessToken Token `json:"access_token"`
	// RefreshToken will be used when access token expired
	// to issue new pair access token and refresh token.
	RefreshToken *Token `json:"refresh_token,omitempty"`
}
