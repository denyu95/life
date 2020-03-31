package service

import (
	"fmt"
	"strconv"
	"time"

	"github.com/denyu95/life/conf"
	"github.com/denyu95/life/models/dao"
	"github.com/denyu95/life/pkg/convertor"
	"github.com/denyu95/life/pkg/db"
	"github.com/denyu95/life/pkg/qq/event"
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
		Uid:      p.Uid,
		Money:    money,
		Remark:   remark,
		CreateAt: p.TimeNow,
		UpdateAt: p.TimeNow,
	}

	if err := spendRecord.Add(); err != nil {
		p.Logger.Error(err)
		replyMsg = conf.SpendFailed
	}

	return replyMsg
}

// 获取一段时间的消费记录
func ListSomeTimeSpendRecord(p *event.ReqParam) (replyMsg string) {
	replyMsg = conf.GetSpendRecordsSuccess
	startDate := ""
	endDate := ""
	if len(p.RegexResult) == 3 {
		startDate = convertor.ToString(p.RegexResult[1])
		endDate = convertor.ToString(p.RegexResult[2])
	} else {
		p.Logger.Error(conf.RegexError)
		replyMsg = conf.GetSpendRecordsFailed
		return
	}

	spendRecord := dao.SpendRecord{}
	var spendRecords []dao.SpendRecord
	if startDate != "" && endDate != "" {
		// 查询月份区间
	} else if startDate != "" {
		// 查询指定月份
		firstOfMonth, _ := time.Parse("2006-01-02 15:04:05", startDate+"-01 00:00:00")
		lastOfMonth := firstOfMonth.AddDate(0, 1, -1)

		strFirstOfMonth := firstOfMonth.Format("2006-01-02") + " 00:00:00"
		strLastOfMonth := lastOfMonth.Format("2006-01-02") + " 23:59:59"
		spendRecords, _ = spendRecord.GetRecordsByConds(map[string]map[string]interface{}{
			"createAt": {">=": strFirstOfMonth, "<=": strLastOfMonth},
		}, "uid ASC, createAt ASC")

		outStr := "—————————————\n"
		out := make(map[string][]dao.SpendRecord, 0)
		for _, spendRecord := range spendRecords {
			v, ok := out[spendRecord.Uid]
			if !ok {
				xx := make([]dao.SpendRecord, 0)
				xx = append(xx, spendRecord)
				out[spendRecord.Uid] = xx
			} else {
				v = append(v, spendRecord)
				out[spendRecord.Uid] = v
			}
		}

		var allTotalMoney float32
		for k, v := range out {
			user := dao.User{}

			user.GetRecordByConds(map[string]map[string]interface{}{
				"uid": {"=": k},
			}, "")
			outStr += user.Name + "\n"
			var totalMoney float32
			for _, vv := range v {
				outStr += vv.CreateAt.Format("06-01-02") + " "
				vMoney := strconv.FormatFloat(float64(vv.Money), 'f', 2, 32)

				o := ""
				st := []rune(vv.Remark)
				for i, v := range st {
					if i == 4 && len(st) > 5 {
						o += "…"
						break
					}
					o += string(v)
				}
				if len(st) < 5 {
					l := 5 - len(st)
					for i := 0; i < l; i++ {
						o += "　"
					}
				}
				vv.Remark = o

				outStr += vv.Remark + " " + vMoney + "\n"
				totalMoney += vv.Money
			}
			allTotalMoney += totalMoney
			outStr += "　　　　　　　总计：" + strconv.FormatFloat(float64(totalMoney), 'f', 2, 32) + "\n"
			outStr += "—————————————\n"
		}

		outStr += "　　　　　全员总计：" + strconv.FormatFloat(float64(allTotalMoney), 'f', 2, 32)

		replyMsg = fmt.Sprintf(replyMsg, "< "+startDate+" > ", outStr)
	} else {
		// 查询本月
		now := time.Now()
		currentYear, currentMonth, _ := now.Date()
		currentLocation := now.Location()

		firstOfMonth := time.Date(currentYear, currentMonth, 1, 0, 0, 0, 0, currentLocation)
		lastOfMonth := firstOfMonth.AddDate(0, 1, -1)

		strFirstOfMonth := firstOfMonth.Format("2006-01-02") + " 00:00:00"
		strLastOfMonth := lastOfMonth.Format("2006-01-02") + " 23:59:59"
		spendRecords, _ = spendRecord.GetRecordsByConds(map[string]map[string]interface{}{
			"createAt": {">=": strFirstOfMonth, "<=": strLastOfMonth},
		}, "uid ASC, createAt ASC")

		outStr := "—————————————\n"
		out := make(map[string][]dao.SpendRecord, 0)
		for _, spendRecord := range spendRecords {
			v, ok := out[spendRecord.Uid]
			if !ok {
				xx := make([]dao.SpendRecord, 0)
				xx = append(xx, spendRecord)
				out[spendRecord.Uid] = xx
			} else {
				v = append(v, spendRecord)
				out[spendRecord.Uid] = v
			}
		}

		var allTotalMoney float32
		for k, v := range out {
			user := dao.User{}

			user.GetRecordByConds(map[string]map[string]interface{}{
				"uid": {"=": k},
			}, "")
			outStr += user.Name + "\n"
			var totalMoney float32
			for _, vv := range v {
				outStr += vv.CreateAt.Format("06-01-02") + " "
				vMoney := strconv.FormatFloat(float64(vv.Money), 'f', 2, 32)

				o := ""
				st := []rune(vv.Remark)
				for i, v := range st {
					if i == 4 && len(st) > 5 {
						o += "…"
						break
					}
					o += string(v)
				}
				if len(st) < 5 {
					l := 5 - len(st)
					for i := 0; i < l; i++ {
						o += "　"
					}
				}
				vv.Remark = o

				outStr += vv.Remark + " " + vMoney + "\n"
				totalMoney += vv.Money
			}
			allTotalMoney += totalMoney
			outStr += "　　　　　　　总计：" + strconv.FormatFloat(float64(totalMoney), 'f', 2, 32) + "\n"
			outStr += "—————————————\n"
		}

		outStr += "　　　　　全员总计：" + strconv.FormatFloat(float64(allTotalMoney), 'f', 2, 32) + "\n"

		// 算余额
		var totalMoney float32
		err := db.GetDB().Raw("SELECT SUM(money) AS totalMoney FROM t_deposit_record").Row().Scan(&totalMoney)
		if err != nil {
			p.Logger.Error(err)
			replyMsg = conf.GetSpendRecordsFailed
			return
		}

		var totalSpendMoney float32
		err = db.GetDB().Raw("SELECT SUM(money) AS totalSpendMoney FROM t_spend_record").Row().Scan(&totalSpendMoney)
		if err != nil {
			p.Logger.Error(err)
			replyMsg = conf.GetSpendRecordsFailed
			return
		}
		remain := totalMoney - totalSpendMoney

		outStr += "　　　　　余　　额：" + strconv.FormatFloat(float64(remain), 'f', 2, 32)

		replyMsg = fmt.Sprintf(replyMsg, "<本月> ", outStr)
	}
	return replyMsg
}
