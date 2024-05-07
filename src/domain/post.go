package domain

import "time"

type Post struct {
	Id        string     `json:"id" bson:"_id,omitempty"`
	Likes     []string   `json:"likes" bson:"likes"`
	Comments  []string   `json:"comments" bson:"comments"`
	Title     string     `json:"title" bson:"title"`
	Channel   Channel    `json:"channel" bson:"channel"`
	Author    User       `json:"author" bson:"author"`
	CreatedAt *time.Time `json:"created_at" bson:"createdAt"`
	UpdatedAt *time.Time `json:"updated_at" bson:"updatedAt"`
}

func (Post) TableName() string { return "posts" }
