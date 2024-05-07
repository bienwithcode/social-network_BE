package domain

import "time"

type Channel struct {
	Id           string     `json:"id" bson:"_id,omitempty"`
	AuthRequired bool       `json:"authRequired" bson:"authRequired"`
	Posts        []*Post    `json:"posts" bson:"posts"`
	Name         string     `json:"name" bson:"name"`
	Order        int32      `json:"order" bson:"order"`
	Description  string     `json:"description" bson:"description"`
	CreatedAt    *time.Time `json:"created_at" bson:"createdAt"`
	UpdatedAt    *time.Time `json:"updated_at" bson:"updatedAt"`
}

func (Channel) TableName() string { return "channels" }
