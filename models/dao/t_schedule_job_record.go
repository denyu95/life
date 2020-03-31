package dao

import "time"

type ScheduleJobRecord struct {
	DaoBase
	Id            int       `gorm:"column:id;primary_key"`                     // 主键id
	Uid           string    `gorm:"column:uid"`                                // 用户id（qq号）
	ScheduleJobId int       `gorm:"column:scheduleJobId" json:"scheduleJobId"` // schedule_job主键
	StartAt       time.Time `gorm:"column:startAt" json:"startAt"`             // 用户确定任务开始时间
	EndAt         time.Time `gorm:"column:endAt" json:"endAt"`                 // 用户确定任务结束时间
	FinishStatus  bool      `gorm:"column:finishStatus" json:"finishStatus"`   // 完成状态 1：完成，0：未完成
}

func (ScheduleJobRecord) TableName() string {
	return "t_schedule_job_record"
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
