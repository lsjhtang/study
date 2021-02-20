package aggregates

import (
	"ddd/domain/models"
	"ddd/domain/repositorys"
)

type Member struct {
	User       *models.User
	Log        *models.UserLog
	userRep    repositorys.IUser
	userLogRep repositorys.IUserLog
}

func NewMember(user *models.User, userRep repositorys.IUser, userLogRep repositorys.IUserLog) *Member {
	return &Member{User: user, userRep: userRep, userLogRep: userLogRep}
}

func NewMemberByName(name string, userRep repositorys.IUser, userLogRep repositorys.IUserLog) *Member {
	user := userRep.FindByName(name)
	return &Member{User: user, userRep: userRep, userLogRep: userLogRep}
}

func (m *Member) Create() error {
	err := m.userRep.SaveUser(m.User)
	if err != nil {
		return err
	}
	m.Log = models.NewUserLog().Set(map[string]interface{}{"LogType": models.UserLogCreate, "LogComment": "创建用户名" + m.Log.UserName})
	return m.userLogRep.SaveUserLog(m.Log)
}
