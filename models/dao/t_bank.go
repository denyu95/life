package dao

import "time"

type Bank struct {
	Id       int       `gorm:"column:id"`
	Uid      string    `gorm:"column:uid"`
	Money    float32   `gorm:"column:money"`
	CreateAt time.Time `gorm:"column:createAt"`
	UpdateAt time.Time `gorm:"column:updateAt"`
}
