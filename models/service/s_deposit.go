package service

import (
	"github.com/denyu95/life/conf"
	"github.com/denyu95/life/models/dao"
	"github.com/denyu95/life/pkg/convertor"
	"github.com/denyu95/life/pkg/qq/event"
)

// 保存充值记录
func SaveDepositRecord(p *event.ReqParam) (replyMsg string) {

	replyMsg = conf.DepositSuccess

	var money float32
	if len(p.RegexResult) == 2 {
		var err error
		money, err = convertor.ToFloat32(p.RegexResult[1])
		if err != nil {
			p.Logger.Error(err)
			replyMsg = conf.DepositFailed
			return
		}
	} else {
		p.Logger.Error(conf.RegexError)
		replyMsg = conf.DepositFailed
		return
	}

	depositRecord := dao.DepositRecord{
		Uid:      p.Uid,
		Money:    money,
		CreateAt: p.TimeNow,
		UpdateAt: p.TimeNow,
	}

	if err := depositRecord.Add(); err != nil {
		p.Logger.Error(err)
		replyMsg = conf.DepositFailed
	}

	return
}
