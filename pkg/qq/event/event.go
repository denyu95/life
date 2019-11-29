package event

import (
	"encoding/json"
	"reflect"
	"regexp"
	"runtime"
	"strings"

	"github.com/sirupsen/logrus"
)

type QQEvent struct {
	privateMsg *PrivateMsg
}

func NewQQEvent() *QQEvent {
	return &QQEvent{}
}

type PrivateMsg struct {
	Font        float64 `json:"font"`
	Message     string  `json:"message"`
	MessageId   int     `json:"message_id"`
	MessageType string  `json:"message_type"`
	PostType    string  `json:"post_type"`
	RawMessage  string  `json:"raw_message"`
	SelfId      float64 `json:"self_id"`
	Sender      Sender  `json:"sender"`
	SubType     string  `json:"sub_type"`
	Time        float64 `json:"time"`
	UserId      float64 `json:"user_id"`
}

type Sender struct {
	Age      int     `json:"age"`
	Nickname string  `json:"nickname"`
	Sex      string  `json:"sex"`
	UserId   float64 `json:"user_id"`
}

type PrivateMsgEvent interface {
	do(msg PrivateMsg)
}

type callPrivateMsgEvent func(p PrivateMsg)

func (pe callPrivateMsgEvent) do(p PrivateMsg) {
	pe(p)
	logrus.WithFields(map[string]interface{}{
		"input":  p,
		"method": runtime.FuncForPC(reflect.ValueOf(pe).Pointer()).Name(),
	}).Info("success")
}

// 提供外部调用
func (qqEvent *QQEvent) OnPrivateMsgEvent(param map[string]interface{}, regex string, f func(PrivateMsg)) {
	if qqEvent.privateMsg == nil {
		qqEvent.privateMsg = new(PrivateMsg)
		//err := mapstructure.Decode(param, qqEvent.privateMsg)
		paramJson, _ := json.Marshal(param)
		err := json.NewDecoder(strings.NewReader(string(paramJson))).Decode(qqEvent.privateMsg)
		if err != nil {
			logrus.Warn(err)
		}
	}
	if ok, _ := regexp.Match(regex, []byte(qqEvent.privateMsg.Message)); ok {
		callPrivateMsgEvent(f).do(*qqEvent.privateMsg)
	}
}
