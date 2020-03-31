package routes

import (
	"github.com/denyu95/life/models/service"
	"github.com/denyu95/life/pkg/qq/event"
)

func HandlePrivateMsg(param map[string]interface{}) {
	event := event.NewQQEvent()
	event.OnPrivateMsgEvent(param, `^加入(?:,|，)([^\n]+)$`, service.SaveUser)
	event.OnPrivateMsgEvent(param, `^hello$`, service.SayHello)

	event.OnPrivateMsgEvent(param, `^充值(?:,|，)(-?\d+\.?\d{0,2})$`, service.SaveDepositRecord)

	// 记录日常消费《！消费备注，消费金额保留两位小数》
	event.OnPrivateMsgEvent(param, `^(?:!|！)([^\n]+)(?:,|，)(\d+\.?\d{0,2})$`, service.SaveSpendRecord)

	// 查看消费清单（不填日期默认查询本月消费）《清单》《清单，2020-01》《清单，2020-01，2020-02》
	event.OnPrivateMsgEvent(param, `^清单(?:(?:,|，)(\d{4}-\d{2}))?(?:(?:,|，)(\d{4}-\d{2}))?$`, service.ListSomeTimeSpendRecord)

	// 新增目标
	event.OnPrivateMsgEvent(param, `^目标(?:,|，)([^\n]+)(?:,|，)(\d+)$`, service.SaveScheduleJob)
	// 目标完成记录
	event.OnPrivateMsgEvent(param, `^收到(?:,|，)(\d{0,9})$`, service.SaveScheduleRecord)
	// 更新目标完成记录
	event.OnPrivateMsgEvent(param, `^完成(?:,|，)(\d{0,9})(?:,|，)(\d)$`, service.UpdateScheduleRecord)
}

func HandleGroupMsg(param map[string]interface{}) {
	event := event.NewQQEvent()
	event.OnGroupMsgEvent(param, `^加入(?:,|，)([^\n]+)$`, service.SaveUser)
	event.OnGroupMsgEvent(param, `^hello$`, service.SayHello)

	event.OnGroupMsgEvent(param, `^充值(?:,|，)(-?\d+\.?\d{0,2})$`, service.SaveDepositRecord)

	// 记录日常消费《！消费备注，消费金额保留两位小数》
	event.OnGroupMsgEvent(param, `^(?:！|!)([^\n]+)(?:,|，)(\d+\.?\d{0,2})$`, service.SaveSpendRecord)

	// 查看消费清单（不填日期默认查询本月消费）《清单》《清单，2020-01》《清单，2020-01，2020-02》
	event.OnGroupMsgEvent(param, `^清单(?:(?:,|，)(\d{4}-\d{2}))?(?:(?:,|，)(\d{4}-\d{2}))?$`, service.ListSomeTimeSpendRecord)
}