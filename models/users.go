package models

import (
    "time"
)

type User struct {
    ID       int64     `gorm:"column:id;primary_key"`
    Name     string    `gorm:"column:name"`
    CreateAt time.Time `gorm:"column:create_at"`
}

func (User) TableName() string {
    return "users"
}
