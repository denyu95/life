package dao

import (
	"time"
)

type SpendRecord struct {
	DaoBase
	Id       int       `gorm:"column:id;primary_key"` // 主键id
	Uid      string    `gorm:"column:uid"`            // 用户id（qq号）
	Money    float32   `gorm:"column:money"`          // 消费金额
	CreateAt time.Time `gorm:"column:createAt"`       // 创建时间
	UpdateAt time.Time `gorm:"column:updateAt"`       // 最后更新时间
}

func (SpendRecord) TableName() string {
	return "t_spend_record"
}
