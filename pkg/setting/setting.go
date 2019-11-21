package setting

import (
	"fmt"
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
)

func init() {
	ini, err := ini.NewIni("conf/app.ini")
	if err != nil {
		fmt.Println(err)
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
}
