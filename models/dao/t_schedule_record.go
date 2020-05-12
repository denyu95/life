package dao

import "time"

type ScheduleJobRecord struct {
	DaoBase
	Id         int       `gorm:"column:id;primary_key"`               // 主键id
	Uid        string    `gorm:"column:uid"`                          // 用户id（qq号）
	ScheduleId int       `gorm:"column:scheduleId" json:"scheduleId"` // schedule_job主键
	ReceiveAt  time.Time `gorm:"column:receiveAt" json:"receiveAt"`   // 收到提醒时间
	IsReceive  bool      `gorm:"column:isReceive" json:"isReceive"`   // 是否收到 1：收到，0：未收到
}

func (ScheduleJobRecord) TableName() string {
	return "t_schedule_record"
}

func (u *ScheduleJobRecord) Update() error {
	return u.DaoBase.Update(u)
}

func (u *ScheduleJobRecord) Add() error {
	return u.DaoBase.Add(u)
}

func (u *ScheduleJobRecord) GetRecordByConds(conds map[string]map[string]interface{}, order string) error {
	return u.DaoBase.GetRecordByConds(u, conds, order)
}

func (u *ScheduleJobRecord) GetRecordsByConds(conds map[string]map[string]interface{}, order string) ([]ScheduleJobRecord, error) {
	scheduleJobRecords := make([]ScheduleJobRecord, 0)
	err := u.DaoBase.GetRecordsByConds(&scheduleJobRecords, conds, order)
	return scheduleJobRecords, err
}
