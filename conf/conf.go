package conf

import (
	"os"

	"github.com/sirupsen/logrus"

	"github.com/denyu95/life/pkg/ini"
)

var (
	Api struct {
		QQBaseUrl string
	}
	MySql struct {
		Username string
		Password string
		Host     string
		Port     int
		DBName   string
		Timeout  string
	}
	Log struct {
		Path  string
		Level string
	}
	Cron struct {
		Goal string
		Boom string
	}
)

func Init() {
	var i *ini.Ini
	var err error

	goEnv := os.Getenv("GO_ENV")
	if goEnv == "prod" {
		i, err = ini.NewIni("conf/apppro.ini")
	} else {
		i, err = ini.NewIni("conf/app.ini")
	}
	if err != nil {
		logrus.Warn(err)
	}
	// Api配置
	Api.QQBaseUrl = i.String("api", "qq.baseUrl")

	// Mysql配置
	MySql.Username = i.String("mysql", "username")
	MySql.Password = i.String("mysql", "password")
	MySql.Host = i.String("mysql", "host")
	MySql.Port, _ = i.Int("mysql", "port")
	MySql.DBName = i.String("mysql", "db.name")
	MySql.Timeout = i.String("mysql", "timeout")

	// Log配置
	Log.Path = i.String("log", "path")
	Log.Level = i.String("log", "level")

	// Cron配置
	//Cron.Goal = i.String("cron", "goal")
	Cron.Boom = i.String("cron", "boom")
}
