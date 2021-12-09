package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"webtest/conf"
)

var Mysql *gorm.DB

func init() {
	var err error
	url := conf.Config.Url
	Mysql, err = gorm.Open("mysql", url)
	if err != nil {
		panic(err)
	}

	// Mysql.LogMode(true)

	Mysql.DB().SetMaxIdleConns(1)
	Mysql.DB().SetMaxOpenConns(10)

	if err != nil {
		panic(err)
	}
}

func CloseMysql() {
	if Mysql != nil {
		Mysql.Close()
	}
}
