package main

import (
	"fmt"
	"strconv"
	"time"

	"github.com/robfig/cron/v3"
	"github.com/sirupsen/logrus"
)

func main() {
	count := 0

	logger := cron.VerbosePrintfLogger(logrus.StandardLogger())

	c := cron.New(
		cron.WithLogger(logger),
		cron.WithChain(cron.DelayIfStillRunning(logger)),
	)
	jobId, _ := c.AddFunc("@every 5s", func() {
		time.Sleep(15 * time.Second)
		count += 1
		fmt.Println("计数器" + strconv.Itoa(count) + time.Now().String())
	})
	c.Start()
	for {
		time.Sleep(time.Second * 30)
		c.Remove(jobId)
	}
}
