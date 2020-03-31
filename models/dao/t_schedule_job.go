package dao

import (
	"time"

	"github.com/denyu95/life/conf"
	"github.com/denyu95/life/pkg/convertor"
	"github.com/denyu95/life/pkg/qq/api"
	"github.com/denyu95/life/pkg/schedule"
)

type ScheduleJob struct {
	DaoBase
	Id         int       `gorm:"column:id;primary_key"`               // 主键id
	Uid        string    `gorm:"column:uid"`                          // 用户id（qq号）
	ChildJobId int       `gorm:"column:childJobId" json:"childJobId"` // 子任务id
	JobId      int       `gorm:"column:jobId" json:"jobId"`           // 任务id
	Message    string    `gorm:"column:message" json:"message"`       // 消息
	Limit      int       `gorm:"column:limit" json:"limit"`           // 任务次数
	CreateAt   time.Time `gorm:"column:createAt" json:"createAt"`     // 创建时间
	DeleteAt   time.Time `gorm:"column:deleteAt" json:"deleteAt"`     // 删除时间
	FinishAt   time.Time `gorm:"column:finishAt" json:"finishAt"`     // 完成时间
	IsDelete   bool      `gorm:"column:isDelete" json:"isDelete"`     // 是否删除
	IsFinish   bool      `gorm:"column:isFinish" json:"isFinish"`     // 是否完成
}

func InitScheduleJob() {
	schedule.Init()
	sj := ScheduleJob{}
	if scheduleJobs, err := sj.GetRecordsByConds(map[string]map[string]interface{}{
		"isDelete": {"=": false},
		"isFinish": {"=": false},
	}, ""); err == nil {
		c := schedule.GetCron()
		for _, scheduleJob := range scheduleJobs {
			jobId, _ := c.AddJob(conf.Cron.Goal, scheduleJob)
			scheduleJob.JobId = int(jobId)
			scheduleJob.Update()
		}
	}
}

func (s ScheduleJob) Run() {
	msg := "提醒：" + s.Message + "，" + convertor.ToString(s.Id)
	api.SendPrivateMsg(map[string]interface{}{
		"user_id": s.Uid,
		"message": msg,
	})

	// boom，更频繁的提醒
	c := schedule.GetCron()
	childJobId, _ := c.AddFunc(conf.Cron.Boom, func() {
		api.SendPrivateMsg(map[string]interface{}{
			"user_id": s.Uid,
			"message": msg,
		})
	})

	s.GetRecordByConds(map[string]map[string]interface{}{
		"id": {"=": s.Id},
	}, "")
	s.ChildJobId = int(childJobId)
	s.Update()
}

func (ScheduleJob) TableName() string {
	return "t_schedule_job"
}

func (u *ScheduleJob) Add() error {
	return u.DaoBase.Add(u)
}

func (u *ScheduleJob) Update() error {
	return u.DaoBase.Update(u)
}

func (u *ScheduleJob) GetRecordByConds(conds map[string]map[string]interface{}, order string) error {
	return u.DaoBase.GetRecordByConds(u, conds, order)
}

func (u *ScheduleJob) GetRecordsByConds(conds map[string]map[string]interface{}, order string) ([]ScheduleJob, error) {
	schedule := make([]ScheduleJob, 0)
	err := u.DaoBase.GetRecordsByConds(&schedule, conds, order)
	return schedule, err
}
