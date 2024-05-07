package domain

import "time"

type Like struct {
	Id        string     `json:"id" bson:"_id,omitempty"`
	User      *User      `json:"user" bson:"user"`
	Post      *Post      `json:"post" bson:"post"`
	CreatedAt *time.Time `json:"created_at" bson:"createdAt"`
	UpdatedAt *time.Time `json:"updated_at" bson:"updatedAt"`
}

func (Like) TableName() string { return "likes" }
