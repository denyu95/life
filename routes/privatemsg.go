package routes

import (
	"github.com/denyu95/life/models/service"
	"github.com/denyu95/life/pkg/qq/event"
)

func HandlePrivateMsg(param map[string]interface{}) {
	event := event.NewQQEvent()
	event.OnPrivateMsgEvent(param, "[a-z]+", service.Demo)
	event.OnPrivateMsgEvent(param, "[A-Z]+", service.Demo1)
}
