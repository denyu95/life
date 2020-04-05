package service

import (
	"github.com/denyu95/life/conf"
	"github.com/denyu95/life/models/dao"
	"github.com/denyu95/life/pkg/convertor"
	"github.com/denyu95/life/pkg/qq/event"
	"github.com/denyu95/life/pkg/schedule"
)

// 保存任务目标到排期
func SaveScheduleJob(p *event.ReqParam) (replyMsg string) {

	replyMsg = conf.ScheduleJobSuccess

	var limit int
	var message string
	if len(p.RegexResult) == 3 {
		message = convertor.ToString(p.RegexResult[1])
		var err error
		limit, err = convertor.ToInt(p.RegexResult[2])
		if err != nil {
			p.Logger.Error(err)
			replyMsg = conf.ScheduleJobFailed
			return
		}
	} else {
		p.Logger.Error(conf.RegexError)
		replyMsg = conf.ScheduleJobFailed
		return
	}

	sch := dao.ScheduleJob{
		Uid:      p.Uid,
		Message:  message,
		Limit:    limit,
		CreateAt: p.TimeNow,
	}

	if err := sch.Add(); err != nil {
		p.Logger.Error(err)
		replyMsg = conf.ScheduleJobFailed
		return
	}

	c := schedule.GetCron()
	jobId, _ := c.AddJob(conf.Cron.Goal, sch)
	sch.JobId = int(jobId)
	sch.Update()

	return replyMsg
}
