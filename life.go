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
)

const Version = "1.0"

func init() {
	log.Init("", time.Minute, time.Minute*5)
}

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
