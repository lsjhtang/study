package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB
var err error

func init() {
	DB, err = gorm.Open("mysql", "dev:dev!123%@(192.168.30.58:3306)/platform_activity?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	DB.LogMode(true)
}
