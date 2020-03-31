package schedule

import (
	"github.com/robfig/cron/v3"
	"github.com/sirupsen/logrus"
)

var c *cron.Cron

func Init() {
	c = GetCron()
	c.Start()
}

func GetCron() *cron.Cron {
	if c == nil {
		logger := cron.VerbosePrintfLogger(logrus.StandardLogger())
		c = cron.New(
			cron.WithLogger(logger),
		)
	}

	return c
}
