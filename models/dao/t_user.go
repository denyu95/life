package dao

import (
	"time"
)

type User struct {
	DaoBase
	Id       int       `gorm:"column:id;primary_key"` // 主键id
	Uid      string    `gorm:"column:uid"`            // 用户id（qq号）
	Name     string    `gorm:"column:name"`           // 用户名
	Nickname string    `gorm:"column:nickname"`       // 昵称
	Sex      int       `gorm:"column:sex"`            // 性别 1：男性 0：女性
	CreateAt time.Time `gorm:"column:createAt"`       // 创建时间
	UpdateAt time.Time `gorm:"column:updateAt"`       // 最后更新时间
}

func (User) TableName() string {
	return "t_user"
}

func (u *User) Add() error {
	return u.DaoBase.Add(u)
}

func (u *User) GetRecordByConds(conds map[string]interface{}, order string) error {
	return u.DaoBase.GetRecordByConds(u, conds, order)
}
