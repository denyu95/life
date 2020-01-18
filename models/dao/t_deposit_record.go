// 充值记录表
package dao

import (
	"time"
)

type DepositRecord struct {
	DaoBase
	Id       int       `gorm:"column:id;primary_key"` // 主键id
	Uid      string    `gorm:"column:uid"`            // 用户id（qq号）
	Money    float32   `gorm:"column:money"`          // 充值金额
	CreateAt time.Time `gorm:"column:createAt"`       // 创建时间
	UpdateAt time.Time `gorm:"column:updateAt"`       // 最后更新时间
}

func (DepositRecord) TableName() string {
	return "t_deposit_record"
}

func (u *DepositRecord) Add() error {
	return u.DaoBase.Add(u)
}

func (u *DepositRecord) GetRecordByConds(conds map[string]interface{}, order string) error {
	return u.DaoBase.GetRecordByConds(u, conds, order)
}