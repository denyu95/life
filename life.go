package main

import (
	"os"
	"os/signal"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"

	"github.com/denyu95/life/cmd"
	"github.com/denyu95/life/pkg/db"
	"github.com/denyu95/life/pkg/log"
	qqApi "github.com/denyu95/life/pkg/qq/api"
	"github.com/denyu95/life/pkg/setting"
)

const Version = "1.0"

// 工程初始化，初始化顺序不要改变
func init() {
	setting.Init()

	logLevel := logrus.InfoLevel
	if setting.Log.Level == "debug" {
		logLevel = logrus.DebugLevel
	}
	log.Init(setting.Log.Path, time.Hour, time.Hour*24*7, logLevel)
	db.Init()
	qqApi.Init()
}

// 1 ） 获取单个对象的方法用 get 做前缀。
// 2 ） 获取多个对象的方法用 list 做前缀。
// 3 ） 获取统计值的方法用 count 做前缀。
// 4 ） 插入的方法用 save（ 推荐 ） 或 insert 做前缀。
// 5 ） 删除的方法用 remove（ 推荐 ） 或 delete 做前缀。
// 6 ） 修改的方法用 update 做前缀。
func main() {
	c := make(chan os.Signal)
	signal.Notify(c)
	defer func() {
		logrus.Info("...关闭数据库连接...")
		db.GetDB().Close()
		logrus.Info("...结束...")
	}()

	logrus.Info("...开始...")

	go do()
	<-c
	logrus.Warn("...外部强制停止...")
}

func do() {
	app := cli.NewApp()
	app.Name = "life"
	app.Usage = "Provide convenience for life"
	app.Version = Version
	app.Commands = []cli.Command{
		cmd.Coolq,
	}
	app.Run(os.Args)
}
