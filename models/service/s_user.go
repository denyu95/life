package service

import (
	"fmt"
	"strings"

	"github.com/denyu95/life/conf"
	"github.com/denyu95/life/models/dao"
	"github.com/denyu95/life/pkg/qq/event"
)

func SaveUser(msg event.PrivateMsg) (replyMsg string) {

	replyMsg = conf.JoinSuccess

	name := ""
	if len(msg.RegxResult) == 2 {
		name = msg.RegxResult[1]
	} else {
		msg.Logger.Error(conf.RegexError)
		replyMsg = conf.JoinFailed
		return
	}

	sex := 1
	if msg.Sender.Sex != "male" {
		sex = 0
	}

	user := dao.User{
		Uid:      msg.Uid,
		Name:     name,
		Nickname: msg.Sender.Nickname,
		Sex:      sex,
		CreateAt: msg.TimeNow,
		UpdateAt: msg.TimeNow,
	}

	if err := user.Add(); err != nil {
		if strings.Contains(err.Error(), "Error 1062") {
			msg.Logger.Warn(err)
			replyMsg = conf.NoNeedToJoin
		} else {
			msg.Logger.Error(err)
			replyMsg = conf.JoinFailed
		}
	}

	return
}

func SayHello(msg event.PrivateMsg) (replyMsg string) {
	replyMsg = "%s%s，你好(⁎⁍̴̛ᴗ⁍̴̛⁎)"

	user := dao.User{}

	user.GetRecordByConds(map[string]interface{}{
		"uid": msg.Uid,
	}, "")

	sex := "小哥哥"
	if user.Sex == 0 {
		sex = "小姐姐"
	}
	replyMsg = fmt.Sprintf(replyMsg, user.Name, sex)
	return
}
