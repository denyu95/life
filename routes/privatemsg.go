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
}
