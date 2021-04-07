package models

import (
	"api/pkg/logging"
	"api/pkg/setting"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

func SetUp() {
	var (
		error        error
		dataBaseTYpe = setting.DatabaseSetting.Type
		user         = setting.DatabaseSetting.User
		pass         = setting.DatabaseSetting.Password
		host         = setting.DatabaseSetting.Host
		name         = setting.DatabaseSetting.Name
	)

	db, error = gorm.Open(dataBaseTYpe, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True", user, pass, host, name))
	if error != nil {
		logging.Fatal("数据库失败", error)
	}

	//设置表名称的前缀
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return setting.DatabaseSetting.TablePrefix + defaultTableName
	}

	db.SingularTable(true) //设置禁用表名的复数形式
	db.LogMode(true)       //打印日志，本地调试的时候可以打开看执行的sql语句

	db.DB().SetMaxIdleConns(10)  //设置空闲时的最大连接数
	db.DB().SetMaxOpenConns(100) //设置数据库的最大打开连接数
}
