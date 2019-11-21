package main

import (
	"github.com/denyu95/life/cmd"
	"github.com/denyu95/life/pkg/db"
	"github.com/urfave/cli"
	"log"
	"os"
	"os/signal"
)

const Version = "1.0"

func main() {
	c := make(chan os.Signal)
	signal.Notify(c)
	defer func() {
		log.Println("...关闭数据库连接...")
		db.GetDB().Close()
		log.Println("...结束...")
	}()

	log.Println("...开始...")
	go do()
	<-c
	log.Println("...外部强制停止...")
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
