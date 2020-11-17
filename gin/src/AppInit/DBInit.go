package AppInit

import (
	"github.com/jinzhu/gorm"
)

var DB *gorm.DB
var err error
func init()  {
	DB, err = gorm.Open("mysql", "root:root@(172.20.10.13:3306)/book?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	DB.LogMode(true)
}

