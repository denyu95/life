package main

import (
	"github.com/denyu95/life/cmd"
	"github.com/denyu95/life/pkg/db"
	"github.com/denyu95/life/pkg/log"
	"github.com/urfave/cli"
	"os"
	"os/signal"
)

const Version = "1.0"

type Param struct {
	A string `json:"a"`
	B string `json:"b"`
}

func main() {
	c := make(chan os.Signal)
	signal.Notify(c)
	defer func() {
		log.Info("...关闭数据库连接...")
		db.GetDB().Close()
		log.Info("...结束...")
	}()

	log.WithField("key", "value").Info("...开始...")

	go do()
	<-c
	log.Warn("...外部强制停止...")
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
