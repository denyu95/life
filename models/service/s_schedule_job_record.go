package service

import (
	"fmt"
	"time"

	"github.com/robfig/cron/v3"

	"github.com/denyu95/life/conf"
	"github.com/denyu95/life/models/dao"
	"github.com/denyu95/life/pkg/convertor"
	"github.com/denyu95/life/pkg/qq/event"
	"github.com/denyu95/life/pkg/schedule"
)

func SaveScheduleRecord(p *event.ReqParam) (replyMsg string) {
	replyMsg = conf.JobRecordFailed
	todayStart := time.Date(p.TimeNow.Year(), p.TimeNow.Month(), p.TimeNow.Day(), 0, 0, 0, 0, p.TimeNow.Location())

	var scheduleJobId int
	if len(p.RegexResult) == 2 {
		scheduleJobId = convertor.MustInt(p.RegexResult[1])
	} else {
		p.Logger.Error(conf.RegexError)
		return
	}

	schR := dao.ScheduleJobRecord{}
	schR.GetRecordByConds(map[string]map[string]interface{}{
		"scheduleJobId": {"=": scheduleJobId},
		"startAt":       {">": todayStart.Format("2006-01-02 15:04:05")},
	}, "")

	schR.Uid = p.Uid
	schR.ScheduleJobId = scheduleJobId
	schR.FinishStatus = false
	schR.StartAt = p.TimeNow

	if err := schR.Update(); err != nil {
		p.Logger.Error(err)
		return
	}

	// 移除boom
	schJob := dao.ScheduleJob{}
	schJob.GetRecordByConds(map[string]map[string]interface{}{
		"id": {"=": scheduleJobId},
	}, "")
	c := schedule.GetCron()
	c.Remove(cron.EntryID(schJob.ChildJobId))

	replyMsg = conf.JobRecordStart

	return replyMsg
}

func UpdateScheduleRecord(p *event.ReqParam) (replyMsg string) {
	replyMsg = conf.JobRecordFailed

	var scheduleJobId int
	var finishStatus bool
	if len(p.RegexResult) == 3 {
		scheduleJobId = convertor.MustInt(p.RegexResult[1])
		finishStatus = convertor.MustBool(p.RegexResult[2])
	} else {
		p.Logger.Error(conf.RegexError)
		return
	}
	todayStart := time.Date(p.TimeNow.Year(), p.TimeNow.Month(), p.TimeNow.Day(), 0, 0, 0, 0, p.TimeNow.Location())

	schR := dao.ScheduleJobRecord{}
	if err := schR.GetRecordByConds(map[string]map[string]interface{}{
		"scheduleJobId": {"=": scheduleJobId},
		"startAt":       {">": todayStart.Format("2006-01-02 15:04:05")},
	}, ""); err != nil {
		p.Logger.Error(err)
		return
	}

	schR.EndAt = p.TimeNow
	schR.FinishStatus = finishStatus

	if err := schR.Update(); err != nil {
		p.Logger.Error(err)
		return
	}

	// 移除boom
	schJob := dao.ScheduleJob{}
	schJob.GetRecordByConds(map[string]map[string]interface{}{
		"id": {"=": scheduleJobId},
	}, "")
	c := schedule.GetCron()
	c.Remove(cron.EntryID(schJob.ChildJobId))

	finishRecords, _ := schR.GetRecordsByConds(map[string]map[string]interface{}{
		"scheduleJobId": {"=": scheduleJobId},
		"finishStatus":  {"=": 1},
	}, "")
	unFinishRecords, _ := schR.GetRecordsByConds(map[string]map[string]interface{}{
		"scheduleJobId": {"=": scheduleJobId},
		"finishStatus":  {"=": 0},
	}, "")
	if finishStatus == true {
		finishCount := len(finishRecords) - len(unFinishRecords)*7
		if finishCount == schJob.Limit {
			replyMsg = conf.JobFinish
			c.Remove(cron.EntryID(schJob.JobId))
			schJob.IsFinish = true
			schJob.FinishAt = p.TimeNow
			schJob.Update()
		} else {
			replyMsg = fmt.Sprintf(conf.JobRecordSuccess, finishCount, schJob.Limit-finishCount)
		}
	} else {
		replyMsg = fmt.Sprintf(conf.JobRecordOhNo, 7)
	}
	return replyMsg
}
