package service

import (
	"strings"
	"time"

	"github.com/denyu95/life/models/dao"
	"github.com/denyu95/life/pkg/convertor"
	"github.com/denyu95/life/pkg/qq/event"
	"github.com/denyu95/life/pkg/log"
)

func CreateUser(msg event.PrivateMsg) (replyMsg string) {
	replyMsg = "创建用户成功"

	timeNow := time.Now()
	strUserId := convertor.ToString(msg.UserId)

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

	if err := user.Add(&user); err != nil {
		log.MapLog[msg.LogId].Warn(err)
		replyMsg = "创建用户失败"
		return
	}

	log.MapLog[msg.LogId].Info(replyMsg)
	return
}
