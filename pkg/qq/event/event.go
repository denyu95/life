package event

import (
	"encoding/json"
	"reflect"
	"regexp"
	"runtime"
	"strings"

	"github.com/sirupsen/logrus"

	"github.com/satori/go.uuid"

	"github.com/denyu95/life/pkg/qq/api"
)

type QQEvent struct {
	privateMsg *PrivateMsg
}

func NewQQEvent() *QQEvent {
	return &QQEvent{}
}

type PrivateMsg struct {
	RegxResult  []string `json:"regx_result"`
	Font        float64  `json:"font"`
	Message     string   `json:"message"`
	MessageId   int      `json:"message_id"`
	MessageType string   `json:"message_type"`
	PostType    string   `json:"post_type"`
	RawMessage  string   `json:"raw_message"`
	SelfId      float64  `json:"self_id"`
	Sender      Sender   `json:"sender"`
	SubType     string   `json:"sub_type"`
	Time        float64  `json:"time"`
	UserId      float64  `json:"user_id"`
	Logger		*logrus.Entry `json:"-"`
}

type Sender struct {
	Age      int     `json:"age"`
	Nickname string  `json:"nickname"`
	Sex      string  `json:"sex"`
	UserId   float64 `json:"user_id"`
}

type PrivateMsgEvent interface {
	do(msg PrivateMsg) string
}

type callPrivateMsgEvent func(msg PrivateMsg) string

func (callEvent callPrivateMsgEvent) do(privateMsg PrivateMsg) {
	uuid, _ := uuid.NewV4()
	logId := strings.ReplaceAll(uuid.String(), "-", "")

	buff, _ := json.Marshal(privateMsg)
	strInput := string(buff)
	methodName := runtime.FuncForPC(reflect.ValueOf(callEvent).Pointer()).Name()
	requestLogger := logrus.WithFields(logrus.Fields{
		"logId" : logId,
		"input" : strInput,
		"method" : methodName,
	})
	privateMsg.Logger = requestLogger

	strOutput := callEvent(privateMsg)

	api.SendPrivateMsg(map[string]interface{}{
		"user_id": privateMsg.Sender.UserId,
		"message": strOutput,
	})

	requestLogger.WithField("output", strOutput)
	requestLogger.Info("success")
}

// 提供外部调用
func (qqEvent *QQEvent) OnPrivateMsgEvent(param map[string]interface{}, strRegex string, f func(PrivateMsg) string) {
	if qqEvent.privateMsg == nil {
		qqEvent.privateMsg = new(PrivateMsg)
		paramJson, _ := json.Marshal(param)
		err := json.NewDecoder(strings.NewReader(string(paramJson))).Decode(qqEvent.privateMsg)
		if err != nil {
			logrus.Warn(err)
		}
	}
	if ok, _ := regexp.Match(strRegex, []byte(qqEvent.privateMsg.Message)); ok {
		regex := regexp.MustCompile(strRegex)
		regxResult := regex.FindStringSubmatch(qqEvent.privateMsg.Message)
		qqEvent.privateMsg.RegxResult = regxResult
		callPrivateMsgEvent(f).do(*qqEvent.privateMsg)
	}
}
