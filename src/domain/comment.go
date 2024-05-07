package domain

import "time"

type Comment struct {
	Id        string     `json:"id" bson:"_id,omitempty"`
	Comment   string     `json:"comment" bson:"comment"`
	Author    *User      `json:"author" bson:"author"`
	Post      *Post      `json:"post" bson:"post"`
	CreatedAt *time.Time `json:"created_at" bson:"createdAt"`
	UpdatedAt *time.Time `json:"updated_at" bson:"updatedAt"`
}

func (Comment) TableName() string { return "comments" }
