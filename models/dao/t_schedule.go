package dao

import (
	"github.com/robfig/cron/v3"
	"github.com/sirupsen/logrus"
	"time"

	"github.com/denyu95/life/conf"
	"github.com/denyu95/life/pkg/qq/api"
	"github.com/denyu95/life/pkg/schedule"
)

type Schedule struct {
	DaoBase
	Id         int       `gorm:"column:id;primary_key"`               // 主键id
	Uid        string    `gorm:"column:uid"`                          // 用户id（qq号）
	ChildJobId int       `gorm:"column:childJobId" json:"childJobId"` // 子任务id
	JobId      int       `gorm:"column:jobId" json:"jobId"`           // 任务id
	Cron       string    `gorm:"column:cron" json:"cron"`             // cron表达式
	Message    string    `gorm:"column:message" json:"message"`       // 消息
	Limit      int       `gorm:"column:limit" json:"limit"`           // 任务次数
	CreateAt   time.Time `gorm:"column:createAt" json:"createAt"`     // 创建时间
	DeleteAt   time.Time `gorm:"column:deleteAt" json:"deleteAt"`     // 删除时间
	FinishAt   time.Time `gorm:"column:finishAt" json:"finishAt"`     // 完成时间
	IsDelete   bool      `gorm:"column:isDelete" json:"isDelete"`     // 是否删除
	IsFinish   bool      `gorm:"column:isFinish" json:"isFinish"`     // 是否完成
}

func InitSchedule() {
	schedule.Init()
	sj := Schedule{}
	if schedules, err := sj.GetRecordsByConds(map[string]map[string]interface{}{
		"isDelete": {"=": false},
		"isFinish": {"=": false},
	}, ""); err == nil {
		c := schedule.GetCron()
		for _, schedule := range schedules {
			jobId, _ := c.AddJob(schedule.Cron, schedule)
			schedule.JobId = int(jobId)
			_ = schedule.Update()
		}
	}
}

func (s Schedule) Run() {
	msg := "提醒：" + s.Message
	api.SendPrivateMsg(map[string]interface{}{
		"user_id": s.Uid,
		"message": msg,
	})

	// 初始化记录
	timeNow := time.Now()
	todayStart := time.Date(timeNow.Year(), timeNow.Month(), timeNow.Day(), 0, 0, 0, 0, timeNow.Location())
	schR := ScheduleJobRecord{}
	err := schR.GetRecordByConds(map[string]map[string]interface{}{
		"scheduleId": {"=": s.Id},
		"uid":        {"=": s.Uid},
		"receiveAt":  {">": todayStart.Format("2006-01-02 15:04:05")},
	}, "")
	if err != nil {
		logrus.Error(err)
	}
	if schR.Id == 0 {
		schR.ScheduleId = s.Id
		schR.Uid = s.Uid
		schR.IsReceive = false
		err := schR.Add()
		if err != nil {
			logrus.Error(err)
		}
	}

	// boom，更频繁的提醒
	c := schedule.GetCron()
	childJobId, _ := c.AddFunc(conf.Cron.Boom, func() {
		if time.Now().Hour() >= 22 {
			c.Remove(cron.EntryID(s.ChildJobId))
			_ = s.GetRecordByConds(map[string]map[string]interface{}{
				"id": {"=": s.Id},
			}, "")
			s.ChildJobId = 0
			_ = s.Update()
		}
		api.SendPrivateMsg(map[string]interface{}{
			"user_id": s.Uid,
			"message": "再次提醒您：" + s.Message,
		})
	})

	_ = s.GetRecordByConds(map[string]map[string]interface{}{
		"id": {"=": s.Id},
	}, "")
	s.ChildJobId = int(childJobId)
	_ = s.Update()
}

func (Schedule) TableName() string {
	return "t_schedule"
}

func (s *Schedule) Add() error {
	return s.DaoBase.Add(s)
}

func (s *Schedule) Update() error {
	return s.DaoBase.Update(s)
}

func (s *Schedule) GetRecordByConds(conds map[string]map[string]interface{}, order string) error {
	return s.DaoBase.GetRecordByConds(s, conds, order)
}

func (s *Schedule) GetRecordsByConds(conds map[string]map[string]interface{}, order string) ([]Schedule, error) {
	schedules := make([]Schedule, 0)
	err := s.DaoBase.GetRecordsByConds(&schedules, conds, order)
	return schedules, err
}
