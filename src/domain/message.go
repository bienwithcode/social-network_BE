package domain

import "time"

type Message struct {
	Id        string     `json:"id" bson:"_id,omitempty"`
	Seen      bool       `json:"seen" bson:"seen"`
	Message   string     `json:"message" bson:"message"`
	Sender    *User      `json:"sender" bson:"sender"`
	Receiver  *User      `json:"receiver" bson:"receiver"`
	CreatedAt *time.Time `json:"created_at" bson:"createdAt"`
	UpdatedAt *time.Time `json:"updated_at" bson:"updatedAt"`
}

func (Message) TableName() string { return "messages" }
