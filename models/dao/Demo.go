package dao

import (
	"github.com/sirupsen/logrus"

	"github.com/denyu95/life/pkg/db"
)

type Demo struct {
	//gorm.Model
	Name string
}

func (Demo) TableName() string {
	return "test"
}

func (d Demo) AddDemo() {
	db := db.GetDB()
	if err := db.Create(d).Error; err != nil {
		logrus.Warn(err)
	}
}
