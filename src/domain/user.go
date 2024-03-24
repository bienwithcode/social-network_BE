package domain

import "time"

type User struct {
	Id            string     `json:"id" bson:"_id,omitempty"`
	Username      string     `json:"username" bson:"username"`
	Email         string     `json:"email" bson:"email"`
	Role          string     `json:"role" bson:"role"`
	EmailVerified bool       `json:"email_verified" bson:"emailVerified"`
	CreatedAt     *time.Time `json:"created_at" bson:"createdAt"`
	UpdatedAt     *time.Time `json:"updated_at" bson:"updatedAt"`
}

func (User) TableName() string { return "users" }
