package valueObjs

import "reflect"

type UserExtra struct {
	Num       int `gorm:"column:num;type:tinyint(2)"`
	UseNum    int `gorm:"column:use_num;type:smallint(4)"`
	Money     int `gorm:"column:money;type:int"`
	CreatedAt int `gorm:"column:created_at;type:int"`
}

func NewUserExtra() *UserExtra {
	return &UserExtra{}
}

func (u *UserExtra) Set(m map[string]interface{}) {
	s := reflect.ValueOf(u).Elem()
	for k, v := range m {
		if s.FieldByName(k).Kind() == reflect.ValueOf(v).Kind() {
			s.FieldByName(k).Set(reflect.ValueOf(v))
		}
	}
}
