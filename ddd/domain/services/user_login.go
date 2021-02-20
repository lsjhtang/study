package services

import (
	"ddd/domain/repositorys"
	"ddd/infrastructure/utils"
	"errors"
)

type UserLogin struct {
	IUser repositorys.IUser
}

func (ul *UserLogin) UserLogin(name, pwd string) (string, error) {
	user := ul.IUser.FindByName(name)
	if user.ID > 0 {
		if utils.MD5(pwd) != user.Openid {
			return "100400", errors.New("用户密码错误")
		}
		return "100200", nil
	} else {
		return "100404", errors.New("用户不存在")
	}
}
