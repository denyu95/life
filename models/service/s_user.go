package service

import (
	"strings"
	"time"

	"fmt"

	"github.com/denyu95/life/models/dao"
	"github.com/denyu95/life/pkg/convertor"
	"github.com/denyu95/life/pkg/qq/event"
)

func CreateUser(msg event.PrivateMsg) (replyMsg string) {
	replyMsg = "加入成功"

	timeNow := time.Now()

	strUserId := convertor.ToString(msg.Sender.UserId)
	userId := strings.Split(strUserId, ".")[0]

	name := ""
	if len(msg.RegxResult) > 1 {
		name = msg.RegxResult[1]
	}

	user := dao.User{
		Uid:      userId,
		Name:     name,
		Nickname: msg.Sender.Nickname,
		CreateAt: timeNow,
		UpdateAt: timeNow,
	}

	if err := user.Add(); err != nil {
		msg.Logger.Warn(err)
		if strings.Contains(err.Error(), "Error 1062") {
			replyMsg = "已加入"
		} else {
			replyMsg = "加入失败"
		}

		return
	}

	return
}

func SayHello(msg event.PrivateMsg) (replyMsg string) {
	replyMsg = "Hello %s!"

	user := dao.User{}
	strUid := convertor.ToString(msg.Sender.UserId)
	uid := strings.Split(strUid, ".")[0]
	user.GetRecordByConds(map[string]interface{}{
		"uid": uid,
	}, "")

	replyMsg = fmt.Sprintf(replyMsg, user.Name)
	return
}
