package service

import (
	"fmt"
	"github.com/denyu95/life/conf"
	"github.com/denyu95/life/models/dao"
	"github.com/denyu95/life/pkg/convertor"
	"github.com/denyu95/life/pkg/qq/event"
	"github.com/denyu95/life/pkg/schedule"
	"github.com/robfig/cron/v3"
)

// 保存提醒排期
func SaveSchedule(p *event.ReqParam) (replyMsg string) {

	replyMsg = conf.ScheduleSuccess

	var cronHour int
	var message string
	if len(p.RegexResult) == 3 {
		message = convertor.ToString(p.RegexResult[1])
		var err error
		cronHour, err = convertor.ToInt(p.RegexResult[2])
		if err != nil {
			p.Logger.Error(err)
			replyMsg = conf.ScheduleFailed
			return
		}
	} else {
		p.Logger.Error(conf.RegexError)
		replyMsg = conf.ScheduleFailed
		return
	}

	sch := dao.Schedule{
		Uid:      p.Uid,
		Message:  message,
		Limit:    90,
		Cron:     "0 " + convertor.ToString(cronHour) + " * * *",
		CreateAt: p.TimeNow,
	}

	if err := sch.Add(); err != nil {
		p.Logger.Error(err)
		replyMsg = conf.ScheduleFailed
		return
	}

	c := schedule.GetCron()
	jobId, _ := c.AddJob(sch.Cron, sch)
	sch.JobId = int(jobId)
	_ = sch.Update()

	return replyMsg
}

// 查看提醒列表
func ListSchedule(p *event.ReqParam) (replyMsg string) {
	replyMsg = conf.ListSchedule

	sch := dao.Schedule{}
	schedules, _ := sch.GetRecordsByConds(map[string]map[string]interface{}{
		"uid":      {"=": p.Uid},
		"isDelete": {"=": false},
	}, "")

	msg := ""
	for _, schedule := range schedules {
		msg += convertor.ToString(schedule.Id) + " " + schedule.Message + "\n"
	}
	return fmt.Sprintf(replyMsg, msg)
}

// 删除提醒
func RemoveSchedule(p *event.ReqParam) (replyMsg string) {
	replyMsg = conf.RemoveScheduleSuccess

	var id int
	if len(p.RegexResult) == 2 {
		var err error
		id, err = convertor.ToInt(p.RegexResult[1])
		if err != nil {
			p.Logger.Error(err)
			replyMsg = conf.RemoveScheduleFailed
			return
		}
	} else {
		p.Logger.Error(conf.RegexError)
		replyMsg = conf.RemoveScheduleFailed
		return
	}

	sch := dao.Schedule{}
	_ = sch.GetRecordByConds(map[string]map[string]interface{}{
		"id": {"=": id},
	}, "")
	sch.IsDelete = true
	sch.DeleteAt = p.TimeNow
	_ = sch.Update()

	c := schedule.GetCron()
	c.Remove(cron.EntryID(sch.JobId))
	c.Remove(cron.EntryID(sch.ChildJobId))
	return replyMsg
}
