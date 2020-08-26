package event

import (
	"encoding/json"
	"reflect"
	"regexp"
	"runtime"
	"strings"
	"time"

	uuid "github.com/satori/go.uuid"
	"github.com/sirupsen/logrus"

	"github.com/denyu95/life/pkg/convertor"
	"github.com/denyu95/life/pkg/qq/api"
)

type QQEvent struct {
	reqParam   *ReqParam
	privateMsg *PrivateMsg
	groupMsg   *GroupMsg
}

func NewQQEvent() *QQEvent {
	return &QQEvent{}
}

type ReqParam struct {
	Uid         string
	Sex         string
	Nickname    string
	Logger      *logrus.Entry `json:"-"`
	TimeNow     time.Time
	RegexResult []string
}

// 私聊
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

// 群聊
type GroupMsg struct {
	Anonymous   Anonymous `json:"anonymous"`
	Font        float64   `json:"font"`
	GroupId     float64   `json:"group_id"`
	Message     string    `json:"message"`
	MessageId   int       `json:"message_id"`
	MessageType string    `json:"message_type"`
	PostType    string    `json:"post_type"`
	RawMessage  string    `json:"raw_message"`
	SelfId      float64   `json:"self_id"`
	Sender      Sender    `json:"sender"`
	SubType     string    `json:"sub_type"`
	Time        float64   `json:"time"`
	UserId      float64   `json:"user_id"`
}

// 匿名者
type Anonymous struct {
	Id   float64 `json:"id"`   // id
	Flag string  `json:"flag"` // 未知
	Name string  `json:"name"` // 匿名名称
}

// 消息发送者
type Sender struct {
	Age      int     `json:"age"`  // 年龄
	Area     string  `json:"area"` // 地区
	Card     string  `json:"card"`
	Level    string  `json:"level"`
	Nickname string  `json:"nickname"` // 昵称
	Role     string  `json:"role"`
	Sex      string  `json:"sex"` // 性别
	Title    string  `json:"title"`
	UserId   float64 `json:"user_id"` // 用户Id
}

type MsgEvent interface {
	do(ReqParam) string
}

type callMsgEvent func(*ReqParam) string

func (callEvent callMsgEvent) do(reqParam *ReqParam) string {
	uuid := uuid.NewV4()
	logId := strings.ReplaceAll(uuid.String(), "-", "")

	strInput := convertor.ToString(reqParam)
	methodName := runtime.FuncForPC(reflect.ValueOf(callEvent).Pointer()).Name()
	requestLogger := logrus.WithFields(logrus.Fields{
		"logId":  logId,
		"input":  strInput,
		"method": methodName,
	})
	reqParam.Logger = requestLogger

	outputMsg := callEvent(reqParam)

	return outputMsg
}

// 提供外部调用
func (qqEvent *QQEvent) OnPrivateMsgEvent(param map[string]interface{}, strRegex string, f func(*ReqParam) string) {
	if qqEvent.privateMsg == nil {
		qqEvent.privateMsg = new(PrivateMsg)
		paramJson, _ := json.Marshal(param)
		err := json.NewDecoder(strings.NewReader(string(paramJson))).Decode(qqEvent.privateMsg)
		if err != nil {
			logrus.Warn(err)
		}

		qqEvent.reqParam = new(ReqParam)
		// 统一处理uid
		qqEvent.reqParam.Uid = convertor.ToString(qqEvent.privateMsg.Sender.UserId)
		// 统一加入当前时间
		qqEvent.reqParam.TimeNow = time.Now()
		// 统一处理性别
		qqEvent.reqParam.Sex = qqEvent.privateMsg.Sender.Sex
		// 统一处理昵称
		qqEvent.reqParam.Nickname = qqEvent.privateMsg.Sender.Nickname
	}
	if ok, _ := regexp.Match(strRegex, []byte(qqEvent.privateMsg.Message)); ok {
		regex := regexp.MustCompile(strRegex)
		regexResult := regex.FindStringSubmatch(qqEvent.privateMsg.Message)
		qqEvent.reqParam.RegexResult = regexResult
		outputMsg := callMsgEvent(f).do(qqEvent.reqParam)
		logrus.Debug("私聊----")
		api.SendMsg(map[string]interface{}{
			"user_id": qqEvent.reqParam.Uid,
			"message": outputMsg,
		})

		qqEvent.reqParam.Logger.WithField("output", outputMsg)
		qqEvent.reqParam.Logger.Info("success")
	}
}

// 提供外部调用
func (qqEvent *QQEvent) OnGroupMsgEvent(param map[string]interface{}, strRegex string, f func(*ReqParam) string) {
	if qqEvent.groupMsg == nil {
		qqEvent.groupMsg = new(GroupMsg)
		paramJson, _ := json.Marshal(param)
		err := json.NewDecoder(strings.NewReader(string(paramJson))).Decode(qqEvent.groupMsg)
		if err != nil {
			logrus.Warn(err)
		}

		qqEvent.reqParam = new(ReqParam)
		// 统一处理uid
		qqEvent.reqParam.Uid = convertor.ToString(qqEvent.groupMsg.Sender.UserId)
		// 统一加入当前时间
		qqEvent.reqParam.TimeNow = time.Now()
		// 统一处理性别
		qqEvent.reqParam.Sex = qqEvent.groupMsg.Sender.Sex
		// 统一处理昵称
		qqEvent.reqParam.Nickname = qqEvent.groupMsg.Sender.Nickname
	}
	if ok, _ := regexp.Match(strRegex, []byte(qqEvent.groupMsg.Message)); ok {
		regex := regexp.MustCompile(strRegex)
		regexResult := regex.FindStringSubmatch(qqEvent.groupMsg.Message)
		qqEvent.reqParam.RegexResult = regexResult
		outputMsg := callMsgEvent(f).do(qqEvent.reqParam)
		logrus.Debug("群聊----")
		logrus.Debug(qqEvent.groupMsg.GroupId)
		logrus.Debug(qqEvent.reqParam.Uid)
		api.SendMsg(map[string]interface{}{
			"message_type": "group",
			"group_id":     qqEvent.groupMsg.GroupId,
			"user_id":      qqEvent.reqParam.Uid,
			"message":      "[CQ:at,qq=" + qqEvent.reqParam.Uid + "] " + outputMsg,
		})

		qqEvent.reqParam.Logger = qqEvent.reqParam.Logger.WithField("output", outputMsg)
		qqEvent.reqParam.Logger.Info("success")
	}
}
