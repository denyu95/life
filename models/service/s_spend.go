package service

import (
	"github.com/denyu95/life/pkg/qq/event"
	"github.com/denyu95/life/conf"
	"github.com/denyu95/life/pkg/convertor"
	"github.com/denyu95/life/models/dao"
)

// 保存消费记录
func SaveSpendRecord(p *event.ReqParam) (replyMsg string) {

	replyMsg = conf.SpendSuccess

	var remark string
	var money float32
	if len(p.RegexResult) == 3 {
		remark = convertor.ToString(p.RegexResult[1])
		var err error
		money, err = convertor.ToFloat32(p.RegexResult[2])
		if err != nil {
			p.Logger.Error(err)
			replyMsg = conf.SpendFailed
			return
		}
	} else {
		p.Logger.Error(conf.RegexError)
		replyMsg = conf.SpendFailed
		return
	}

	spendRecord := dao.SpendRecord{
		Uid:p.Uid,
		Money:money,
		Remark: remark,
		CreateAt: p.TimeNow,
		UpdateAt: p.TimeNow,
	}

	if err := spendRecord.Add(); err != nil {
		p.Logger.Error(err)
		replyMsg = conf.SpendFailed
	}

	return replyMsg
}
