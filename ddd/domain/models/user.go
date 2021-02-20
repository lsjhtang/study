package models

import (
	"ddd/domain/valueObjs"
	"reflect"
)

type User struct {
	*Model
	ID       int    `gorm:"column:id;AUTO_INCREMENT;PRIMARY_KEY"`
	Openid   string `gorm:"column:openid;type:varchar(200)"`
	UserId   string `gorm:"column:user_id;type:varchar(20)"`
	NickName string `gorm:"column:nick_name;type:varchar(200)"`
	UserName string `gorm:"column:user_name;type:varchar(200)"`
	RoleId   string `gorm:"column:role_id;type:varchar(30)"`
	ServerId string `gorm:"column:server_id;type:varchar(20)"`
	Extra    *valueObjs.UserExtra
}

func NewUser() *User {
	user := &User{}
	user.EntityId = 2
	user.EntityName = "user entity"
	return user
}

func (u *User) Set(m map[string]interface{}) *User {
	s := reflect.ValueOf(u).Elem()
	for k, v := range m {
		if s.FieldByName(k).Kind() == reflect.ValueOf(v).Kind() {
			s.FieldByName(k).Set(reflect.ValueOf(v))
		}
	}

	return u
}
