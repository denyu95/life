package dao

import (
	"time"
)

type User struct {
	DaoBase
	Id       int       `gorm:"column:id;primary_key"`
	Uid      string    `gorm:"column:uid"`
	Name     string    `gorm:"column:name"`
	Nickname string    `gorm:"column:nickname"`
	CreateAt time.Time `gorm:"column:createAt"`
	UpdateAt time.Time `gorm:"column:updateAt"`
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
