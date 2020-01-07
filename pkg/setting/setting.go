package setting

import (
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
		Path string
		Level string
	}
)

func Init() {
	ini, err := ini.NewIni("conf/app.ini")
	if err != nil {
		logrus.Warn(err)
	}
	// Api配置
	Api.QQBaseUrl = ini.String("api", "qq.baseUrl")

	// Mysql配置
	MySql.Username = ini.String("mysql", "username")
	MySql.Password = ini.String("mysql", "password")
	MySql.Host = ini.String("mysql", "host")
	MySql.Port, _ = ini.Int("mysql", "port")
	MySql.DBName = ini.String("mysql", "db.name")
	MySql.Timeout = ini.String("mysql", "timeout")

	// Log配置
	Log.Path = ini.String("log", "path")
	Log.Level = ini.String("log", "level")
}
