package dao

import (
	"time"
)

type SpendRecord struct {
	DaoBase
	Id       int       `gorm:"column:id;primary_key"`
	Uid      string    `gorm:"column:uid"`
	Money    float32   `gorm:"column:money"`
	CreateAt time.Time `gorm:"column:createAt"`
	UpdateAt time.Time `gorm:"column:updateAt"`
}

func (SpendRecord) TableName() string {
	return "t_spend_record"
}
