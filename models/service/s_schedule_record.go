package service

import (
	"github.com/denyu95/life/conf"
	"github.com/denyu95/life/models/dao"
	"github.com/denyu95/life/pkg/qq/event"
	"github.com/denyu95/life/pkg/schedule"
	"github.com/robfig/cron/v3"
)

func UpdateScheduleRecord(p *event.ReqParam) (replyMsg string) {
	replyMsg = conf.JobRecordFailed
	schR := dao.ScheduleJobRecord{}
	if sjrs, err := schR.GetRecordsByConds(map[string]map[string]interface{}{
		"uid":       {"=": p.Uid},
		"receiveAt": {"=": "0000-00-00 00:00:00"},
	}, ""); err != nil {
		p.Logger.Error(err)
	} else {
		if len(sjrs) == 0 {
			replyMsg = conf.JobRecordStopDo
			return
		}
		for _, sjr := range sjrs {
			sjr.IsReceive = true
			sjr.ReceiveAt = p.TimeNow
			_ = sjr.Update()

			sch := dao.Schedule{}
			_ = sch.GetRecordByConds(map[string]map[string]interface{}{
				"id": {"=": sjr.ScheduleId},
			}, "")
			c := schedule.GetCron()
			c.Remove(cron.EntryID(sch.JobId))
			c.Remove(cron.EntryID(sch.ChildJobId))
		}
		replyMsg = conf.JobRecordSuccess
	}
	// TODO 提醒达到限制次数则完成提醒
	return
}
