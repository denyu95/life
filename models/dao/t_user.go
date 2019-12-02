package dao

import (
	"time"
)

type User struct {
	Base
	Id       int       `gorm:"column:id"`
	Uid      string    `gorm:"column:uid"`
	Name     string    `gorm:"column:name"`
	Nickname string    `gorm:"column:nickname"`
	CreateAt time.Time `gorm:"column:createAt"`
	UpdateAt time.Time `gorm:"column:updateAt"`
}

func (User) TableName() string {
	return "t_user"
}
