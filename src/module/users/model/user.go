package usermodel

import "time"

type User struct {
	Id        int        `json:"id" gorm:"column:id;"`
	Username  string     `json:"username" gorm:"column:username;"`
	Status    int        `json:"status" gorm:"column:status;"`
	CreatedAt *time.Time `json:"created_at" gorm:"column:created_at;"`
	UpdatedAt *time.Time `json:"updated_at" gorm:"column:updated_at;"`
}

func (User) TableName() string { return "Users" }
