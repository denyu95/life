package event

import (
	"github.com/mitchellh/mapstructure"
	"log"
	"regexp"
)

type QQEvent struct {
	privateMsg *PrivateMsg
}

func NewQQEvent() *QQEvent {
	return &QQEvent{}
}

type PrivateMsg struct {
	Font        float64 `mapstructure:"font"`
	Message     string  `mapstructure:"message"`
	MessageId   int     `mapstructure:"message_id"`
	MessageType string  `mapstructure:"message_type"`
	PostType    string  `mapstructure:"post_type"`
	RawMessage  string  `mapstructure:"raw_message"`
	SelfId      float64 `mapstructure:"self_id"`
	Sender      Sender  `mapstructure:"sender"`
	SubType     string  `mapstructure:"sub_type"`
	Time        float64 `mapstructure:"time"`
	UserId      float64 `mapstructure:"user_id"`
}

type Sender struct {
	Age      int     `mapstructure:"age"`
	Nickname string  `mapstructure:"nickname"`
	Sex      string  `mapstructure:"sex"`
	UserId   float64 `mapstructure:"user_id"`
}

type PrivateMsgEvent interface {
	do(msg PrivateMsg)
}

type callPrivateMsgEvent func(PrivateMsg)

func (pe callPrivateMsgEvent) do(p PrivateMsg) {
	pe(p)
}

// 提供外部调用
func (qqEvent *QQEvent) OnPrivateMsgEvent(param map[string]interface{}, regex string, f func(PrivateMsg)) {
	if qqEvent.privateMsg == nil {
		qqEvent.privateMsg = new(PrivateMsg)
		err := mapstructure.Decode(param, qqEvent.privateMsg)
		if err != nil {
			log.Println(err)
		}
	}
	if ok, _ := regexp.Match(regex, []byte(qqEvent.privateMsg.Message)); ok {
		callPrivateMsgEvent(f).do(*qqEvent.privateMsg)
	}
}
