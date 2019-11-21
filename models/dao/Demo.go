package dao

import (
	"github.com/denyu95/life/pkg/db"
	"log"
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
		log.Println("插入失败", err)
	}
}
