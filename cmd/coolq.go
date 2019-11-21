package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/denyu95/life/pkg/convert"
	"github.com/denyu95/life/routes"
	"github.com/urfave/cli"
	"io/ioutil"
	"log"
	"net/http"
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
		fmt.Println(err)
	}
	return err
}

// coolq事件回掉接口
func coolqEvent(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	buf, _ := ioutil.ReadAll(r.Body)
	m := map[string]interface{}{}
	json.Unmarshal(buf, &m)
	// 调试专用打印
	fmt.Println(m)

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
		log.Println("QQ未知post请求", strPostType)
	}
}

func msgTypeEvent(msgType string, param map[string]interface{}) {
	if msgType == "private" {
		log.Println("private")
		routes.HandlePrivateMsg(param)
	} else if msgType == "group" {
		log.Println("group")
	} else if msgType == "discuss" {
		log.Println("discuss")
	} else {
		log.Println("QQ未知消息类型")
	}
}

func eventTypeEvent(eventType string) {
	if eventType == "group_upload" {
		log.Println("group_upload")
	} else if eventType == "group_admin" {
		log.Println("group_admin")
	} else if eventType == "group_decrease" {
		log.Println("group_decrease")
	} else if eventType == "group_increase" {
		log.Println("group_increase")
	} else {
		log.Println("QQ未知事件类型")
	}
}

func requestTypeEvent(reqType string) {
	if reqType == "friend" {
		log.Println("friend")
	} else if reqType == "group" {
		log.Println("group")
	} else {
		log.Println("QQ未知请求类型")
	}
}
