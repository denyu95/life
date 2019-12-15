package dao

import (
	"github.com/denyu95/life/pkg/db"
)

type DaoBase struct{}

func (b DaoBase) Add(m interface{}) error {
	db := db.GetDB()
	return db.Create(m).Error
}

func (b *DaoBase) GetRecordByConds(m interface{}, conds map[string]interface{}, order string) error {
	query := ""
	args := make([]interface{}, 0)
	if conds != nil {
		for k, v := range conds {
			query += k + " = ? AND "
			args = append(args, v)
		}
		if query[len(query)-4:len(query)-1] == "AND" {
			query = query[0 : len(query)-4]
		}
	}

	err := db.GetDB().Where(query, args).
		Order(order).
		First(m).
		GetErrors()
	if len(err) > 0 {
		return err[0]
	} else {
		return nil
	}
}
