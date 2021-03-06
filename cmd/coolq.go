package cmd

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"

	"github.com/denyu95/life/pkg/convertor"
	"github.com/denyu95/life/routes"
)

var Coolq = cli.Command{
	Name:        "coolq",
	Usage:       "",
	Description: "",
	Action:      runCoolq,
}

var (
	upgrader = websocket.Upgrader{
		// 允许跨域访问
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
)

func runCoolq(c *cli.Context) error {
	http.HandleFunc("/", coolqEvent)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		logrus.Warn(err)
	}
	return err
}

func coolqEvent(w http.ResponseWriter, r *http.Request) {
	// 收到http请求，协议升级转换成websocket
	// var (
	// 	conn *websocket.Conn
	// 	err  error
	// 	buf  []byte
	// )

	// if conn, err = upgrader.Upgrade(w, r, nil); err != nil {
	// 	// 链接终止
	// 	return
	// }

	// for {
	// if _, buf, err = conn.ReadMessage(); err != nil {
	// 	// 关闭websocket，链接终止
	// 	goto ERR
	// }
	buf, _ := ioutil.ReadAll(r.Body)
	m := map[string]interface{}{}
	json.Unmarshal(buf, &m)

	strPostType := convertor.ToString(m["post_type"])
	strMsgType := convertor.ToString(m["message_type"])
	strEventType := convertor.ToString(m["event"])
	strReqType := convertor.ToString(m["request_type"])

	if strPostType == "message" {
		msgTypeEvent(strMsgType, m)
	} else if strPostType == "event" {
		eventTypeEvent(strEventType)
	} else if strPostType == "request" {
		requestTypeEvent(strReqType)
	} else {
		logrus.Info("QQ未知post请求", strPostType)
	}

	//发送数据，判断返回值是否报错
	// if err = conn.WriteMessage(websocket.TextMessage, []byte("{\"result\":\"ok\"}")); err != nil {
	// 	//报错了
	// 	goto ERR
	// }
	// }
	// error的标签
	// ERR:
	// 	conn.Close()
}

func msgTypeEvent(msgType string, param map[string]interface{}) {
	if msgType == "private" {
		routes.HandlePrivateMsg(param)
	} else if msgType == "group" {
		//logrus.Info("group")
		routes.HandleGroupMsg(param)
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
