package service

import (
	"github.com/denyu95/life/conf"
	"github.com/denyu95/life/models/dao"
	"github.com/denyu95/life/pkg/convertor"
	"github.com/denyu95/life/pkg/qq/event"
)

func SaveDepositRecord(msg event.PrivateMsg) (replyMsg string) {

	replyMsg = conf.DepositSuccess

	var money float32
	if len(msg.RegxResult) == 2 {
		var err error
		money, err = convertor.ToFloat32(msg.RegxResult[1])
		if err != nil {
			msg.Logger.Error(err)
			replyMsg = conf.DepositFailed
			return
		}
	} else {
		msg.Logger.Error(conf.RegexError)
		replyMsg = conf.DepositFailed
		return
	}

	depositRecord := dao.DepositRecord{
		Uid:      msg.Uid,
		Money:    money,
		CreateAt: msg.TimeNow,
		UpdateAt: msg.TimeNow,
	}

	if err := depositRecord.Add(); err != nil {
		msg.Logger.Error(err)
		replyMsg = conf.DepositFailed
	}

	return
}
