package dao

import (
	"time"
)

type DepositRecord struct {
	DaoBase
	Id       int       `gorm:"column:id;primary_key"`
	Uid      string    `gorm:"column:uid"`
	Money    float32   `gorm:"column:money"`
	CreateAt time.Time `gorm:"column:createAt"`
	UpdateAt time.Time `gorm:"column:updateAt"`
}

func (DepositRecord) TableName() string {
	return "t_deposit_record"
}
