package dao

import "time"

type User struct {
	Id       int       `gorm:"column:id"`
	Uid      string    `gorm:"column:uid"`
	Name     string    `gorm:"column:name"`
	CreateAt time.Time `gorm:"column:createAt"`
	UpdateAt time.Time `gorm:"column:updateAt"`
}
