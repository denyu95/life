package cmd

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"

	"github.com/denyu95/life/pkg/convert"
	"github.com/denyu95/life/routes"
)

var Coolq = cli.Command{
	Name:        "coolq",
	Usage:       "",
	Description: "",
	Action:      runCoolq,
}

func runCoolq(c *cli.Context) error {
	http.HandleFunc("/", coolqEvent)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		logrus.Warn(err)
	}
	return err
}

// coolq事件回掉接口
func coolqEvent(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	buf, _ := ioutil.ReadAll(r.Body)
	m := map[string]interface{}{}
	json.Unmarshal(buf, &m)

	strPostType, _ := convert.ToString(m["post_type"])
	strMsgType, _ := convert.ToString(m["message_type"])
	strEventType, _ := convert.ToString(m["event"])
	strReqType, _ := convert.ToString(m["request_type"])

	if strPostType == "message" {
		msgTypeEvent(strMsgType, m)
	} else if strPostType == "event" {
		eventTypeEvent(strEventType)
	} else if strPostType == "request" {
		requestTypeEvent(strReqType)
	} else {
		logrus.Info("QQ未知post请求", strPostType)
	}
}

func msgTypeEvent(msgType string, param map[string]interface{}) {
	if msgType == "private" {
		routes.HandlePrivateMsg(param)
	} else if msgType == "group" {
		logrus.Info("group")
	} else if msgType == "discuss" {
		logrus.Info("discuss")
	} else {
		logrus.Info("QQ未知消息类型")
	}
}

func eventTypeEvent(eventType string) {
	if eventType == "group_upload" {
		logrus.Info("group_upload")
	} else if eventType == "group_admin" {
		logrus.Info("group_admin")
	} else if eventType == "group_decrease" {
		logrus.Info("group_decrease")
	} else if eventType == "group_increase" {
		logrus.Info("group_increase")
	} else {
		logrus.Info("QQ未知事件类型")
	}
}

func requestTypeEvent(reqType string) {
	if reqType == "friend" {
		logrus.Info("friend")
	} else if reqType == "group" {
		logrus.Info("group")
	} else {
		logrus.Info("QQ未知请求类型")
	}
}
