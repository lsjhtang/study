package dao

import (
	Models "ddd/domain/models"
)

type UserLog struct {
}

func newUserLog() *UserLog {
	return &UserLog{}
}

func (ul *UserLog) FindByName(name string) *Models.UserLog {
	return nil
}

func (ul *UserLog) SaveUserLog(*Models.UserLog) error {
	return nil
}
