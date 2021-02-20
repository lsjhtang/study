package models

import "reflect"

const (
	UserLogCreate = 5
	UserLogUpdate = 6
)

type UserLog struct {
	*Model
	ID         int    `gorm:"column:id;AUTO_INCREMENT;PRIMARY_KEY"`
	UserId     string `gorm:"column:user_id;type:varchar(50)"`
	UserName   string `gorm:"column:user_name;type:varchar(255)"`
	LogType    string `gorm:"column:log_type;type:smallint(4)"`
	LogComment string `gorm:"column:log_comment;type:varchar(255)"`
	UpdateAt   string `gorm:"column:update_at;type:int(11)"`
}

func NewUserLog() *UserLog {
	return &UserLog{}
}

func (u *UserLog) Set(m map[string]interface{}) *UserLog {
	s := reflect.ValueOf(u).Elem()
	for k, v := range m {
		if s.FieldByName(k).Kind() == reflect.ValueOf(v).Kind() {
			s.FieldByName(k).Set(reflect.ValueOf(v))
		}
	}
	return u
}
