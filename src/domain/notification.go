package domain

import "time"

type Notification struct {
	Id        string     `json:"id" bson:"_id,omitempty"`
	Seen      bool       `json:"seen" bson:"seen"`
	Author    *User      `json:"author" bson:"author"`
	User      *User      `json:"user" bson:"user"`
	Post      *Post      `json:"post" bson:"post"`
	Message   *Message   `json:"message" bson:"message"`
	CreatedAt *time.Time `json:"created_at" bson:"createdAt"`
	UpdatedAt *time.Time `json:"updated_at" bson:"updatedAt"`
}

func (Notification) TableName() string { return "notifications" }
