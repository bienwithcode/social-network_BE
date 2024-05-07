package domain

import "time"

type Follow struct {
	Id        string     `json:"id" bson:"_id,omitempty"`
	User      *User      `json:"user" bson:"user"`
	Follower  *User      `json:"follower" bson:"follower"`
	CreatedAt *time.Time `json:"created_at" bson:"createdAt"`
	UpdatedAt *time.Time `json:"updated_at" bson:"updatedAt"`
}

func (Follow) TableName() string { return "follows" }
