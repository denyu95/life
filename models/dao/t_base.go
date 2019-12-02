package dao

import (
	"github.com/denyu95/life/pkg/db"
)

type Base struct {
}

func (b Base) Add(m interface{}) error {
	db := db.GetDB()
	return db.Create(m).Error
}
