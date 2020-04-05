package db

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/sirupsen/logrus"

	"github.com/denyu95/life/conf"
	"time"
)

var _db *gorm.DB

func Init() {
	username := conf.MySql.Username //账号
	password := conf.MySql.Password //密码
	host := conf.MySql.Host         //数据库地址，可以是Ip或者域名
	port := conf.MySql.Port         //数据库端口
	dbName := conf.MySql.DBName     //数据库名
	timeout := conf.MySql.Timeout   //连接超时，10秒

	// 拼接下dsn参数, dsn格式可以参考上面的语法，这里使用Sprintf动态拼接dsn参数，因为一般数据库连接参数，我们都是保存在配置文件里面，需要从配置文件加载参数，然后拼接dsn。
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local&timeout=%s", username, password, host, port, dbName, timeout)
	var err error
	_db, err = gorm.Open("mysql", dsn)
	if err != nil {
		logrus.Error(err)
	}
	// 设置数据库连接池参数
	_db.DB().SetMaxOpenConns(100) // 数据库连接池最大连接数
	_db.DB().SetMaxIdleConns(20)  // 连接池空闲连接数，如果最大连接数设置小于空闲连接数，则空闲连接数等于最大连接数。
	// 解决invalid connection
	// 现象：服务长时间未被请求就会出现invalid connection，连接数据库出问题
	// mysql 连接池中的连接被单方面关闭了，而程序却不知道，依然使用这个连接，所以会出现这个错误
	// 解决办法：为客户端的连接池设置一个更短的生存时间。
	_db.DB().SetConnMaxLifetime(60 * time.Second)
	// 关闭调试日志
	_db.LogMode(false)
	// 开启调试日志
	//_db.LogMode(true)
}

func GetDB() *gorm.DB {
	if _db.Error != nil {
		logrus.Error(_db.Error)
	}
	return _db
}
