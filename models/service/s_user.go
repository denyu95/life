package service

import (
	"fmt"
	"strings"

	"github.com/denyu95/life/conf"
	"github.com/denyu95/life/models/dao"
	"github.com/denyu95/life/pkg/convertor"
	"github.com/denyu95/life/pkg/qq/api"
	"github.com/denyu95/life/pkg/qq/event"
)

// 保存用户
func SaveUser(p *event.ReqParam) (replyMsg string) {

	replyMsg = conf.JoinSuccess

	name := ""
	if len(p.RegexResult) == 2 {
		name = p.RegexResult[1]
	} else {
		p.Logger.Error(conf.RegexError)
		replyMsg = conf.JoinFailed
		return
	}

	sex := 1
	if p.Sex == "unknown" {
		userInfo := api.GetStrangerInfo(map[string]interface{}{"user_id": p.Uid})
		p.Sex = convertor.ToString(userInfo["sex"])
	}
	if p.Sex != "male" {
		sex = 0
	}

	user := dao.User{
		Uid:      p.Uid,
		Name:     name,
		Nickname: p.Nickname,
		Sex:      sex,
		CreateAt: p.TimeNow,
		UpdateAt: p.TimeNow,
	}

	if err := user.Add(); err != nil {
		if strings.Contains(err.Error(), "Error 1062") {
			p.Logger.Warn(err)
			replyMsg = conf.NoNeedToJoin
		} else {
			p.Logger.Error(err)
			replyMsg = conf.JoinFailed
		}
	}

	return
}

// 打招呼
func SayHello(p *event.ReqParam) (replyMsg string) {
	replyMsg = "%s%s，你好(⁎⁍̴̛ᴗ⁍̴̛⁎)"

	user := dao.User{}

	user.GetRecordByConds(map[string]interface{}{
		"uid": p.Uid,
	}, "")

	sex := "小哥哥"
	if user.Sex == 0 {
		sex = "小姐姐"
	}
	replyMsg = fmt.Sprintf(replyMsg, user.Name, sex)
	return
}
