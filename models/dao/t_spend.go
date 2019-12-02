package dao

import (
	"time"
)

type Spend struct {
	Base
	Id       int       `gorm:"column:id"`
	Uid      string    `gorm:"column:uid"`
	Money    float32   `gorm:"column:money"`
	CreateAt time.Time `gorm:"column:createAt"`
	UpdateAt time.Time `gorm:"column:updateAt"`
}

func (Spend) TableName() string {
	return "t_spend"
}
