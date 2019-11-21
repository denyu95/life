package db

import (
	"fmt"
	"github.com/denyu95/life/pkg/setting"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var _db *gorm.DB

func init() {
	username := setting.MySql.Username //账号
	password := setting.MySql.Password //密码
	host := setting.MySql.Host         //数据库地址，可以是Ip或者域名
	port := setting.MySql.Port         //数据库端口
	dbName := setting.MySql.DBName     //数据库名
	timeout := setting.MySql.Timeout   //连接超时，10秒

	// 拼接下dsn参数, dsn格式可以参考上面的语法，这里使用Sprintf动态拼接dsn参数，因为一般数据库连接参数，我们都是保存在配置文件里面，需要从配置文件加载参数，然后拼接dsn。
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local&timeout=%s", username, password, host, port, dbName, timeout)
	var err error
	_db, err = gorm.Open("mysql", dsn)
	if err != nil {
		panic("连接数据库失败，error=" + err.Error())
	}
	// 设置数据库连接池参数
	_db.DB().SetMaxOpenConns(100) // 数据库连接池最大连接数
	_db.DB().SetMaxIdleConns(20)  // 连接池空闲连接数，如果最大连接数设置小于空闲连接数，则空闲连接数等于最大连接数。
}

func GetDB() *gorm.DB {
	return _db
}
