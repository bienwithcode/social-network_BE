package model

import "social-network/domain"

type AuthUserId struct {
	Id string `json:"id" form:"_id"`
}

type AuthEmailPassword struct {
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

type User struct {
	Id    string `json:"id" bson:"_id,omitempty"`
	Email string `json:"email" bson:"email"`
	Role  string `json:"role" bson:"role"`
}

type TokenResponse struct {
	Token string `json:"token"`
	// RefreshToken will be used when access token expired
	// to issue new pair access token and refresh token.
	// RefreshToken *Token `json:"refresh_token,omitempty"`
	User *domain.User `json:"user"`
}

type Filter struct {
	EmailVerified   *bool   `json:"emailVerified,omitempty" form:"emailVerified"`
	HideBannedUsers *bool   `json:"hideBannedUsers,omitempty" form:"hideBannedUsers"`
	SearchQuery     *string `json:"searchQuery,omitempty" form:"searchQuery"`
}
