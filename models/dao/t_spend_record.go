// 消费记录表
package dao

import (
	"time"
)

type SpendRecord struct {
	DaoBase
	Id       int       `gorm:"column:id;primary_key"` // 主键id
	Uid      string    `gorm:"column:uid"`            // 用户id（qq号）
	Money    float32   `gorm:"column:money"`          // 消费金额
	Remark   string    `gorm:"remark"`                // 消费备注
	CreateAt time.Time `gorm:"column:createAt"`       // 创建时间
	UpdateAt time.Time `gorm:"column:updateAt"`       // 最后更新时间
}

func (SpendRecord) TableName() string {
	return "t_spend_record"
}

func (u *SpendRecord) Add() error {
	return u.DaoBase.Add(u)
}

func (u *SpendRecord) GetRecordByConds(conds map[string]interface{}, order string) error {
	return u.DaoBase.GetRecordByConds(u, conds, order)
}

func (u *SpendRecord) GetRecordsByConds(conds map[string]map[string]interface{}, order string) ([]SpendRecord, error) {
	spendRecords := make([]SpendRecord, 0)
	err := u.DaoBase.GetRecordsByConds(&spendRecords, conds, order)
	return spendRecords, err
}
