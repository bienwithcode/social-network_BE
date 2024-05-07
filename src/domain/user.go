package domain

import "time"

type User struct {
	Id            string     `json:"id" bson:"_id,omitempty"`
	Role          string     `json:"role" bson:"role"`
	EmailVerified bool       `json:"email_verified" bson:"emailVerified"`
	Banned        bool       `json:"banned" bson:"banned"`
	FacebookId    string     `json:"facebookId" bson:"facebookId"`
	GoogleId      string     `json:"googleId" bson:"googleId"`
	GithubId      string     `json:"githubId" bson:"githubId"`
	IsOnline      bool       `json:"isOnline" bson:"isOnline"`
	Posts         []string   `json:"posts" bson:"posts"`
	Likes         []string   `json:"likes" bson:"likes"`
	Comments      []string   `json:"comments" bson:"comments"`
	Followers     []string   `json:"followers" bson:"followers"`
	Following     []string   `json:"following" bson:"following"`
	Messages      []string   `json:"messages" bson:"messages"`
	Notifications []string   `json:"notifications" bson:"notifications"`
	FullName      string     `json:"fullName" bson:"fullName"`
	Email         string     `json:"email" bson:"email"`
	CreatedAt     *time.Time `json:"created_at" bson:"createdAt"`
	UpdatedAt     *time.Time `json:"updated_at" bson:"updatedAt"`
}

type UserWithPassword struct {
	Id            string     `json:"id" bson:"_id,omitempty"`
	Role          string     `json:"role" bson:"role"`
	EmailVerified bool       `json:"email_verified" bson:"emailVerified"`
	Banned        bool       `json:"banned" bson:"banned"`
	FacebookId    string     `json:"facebookId" bson:"facebookId"`
	GoogleId      string     `json:"googleId" bson:"googleId"`
	GithubId      string     `json:"githubId" bson:"githubId"`
	IsOnline      bool       `json:"isOnline" bson:"isOnline"`
	Posts         []string   `json:"posts" bson:"posts"`
	Likes         []string   `json:"likes" bson:"likes"`
	Comments      []string   `json:"comments" bson:"comments"`
	Followers     []string   `json:"followers" bson:"followers"`
	Following     []string   `json:"following" bson:"following"`
	Messages      []string   `json:"messages" bson:"messages"`
	Notifications []string   `json:"notifications" bson:"notifications"`
	FullName      string     `json:"fullName" bson:"fullName"`
	Email         string     `json:"email" bson:"email"`
	CreatedAt     *time.Time `json:"created_at" bson:"createdAt"`
	UpdatedAt     *time.Time `json:"updated_at" bson:"updatedAt"`
	Password      string     `json:"password" bson:"password"`
}

func (User) TableName() string { return "users" }
