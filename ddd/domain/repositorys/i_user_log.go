package repositorys

import (
	Models "ddd/domain/models"
)

type IUserLog interface {
	FindByName(name string) *Models.UserLog
	SaveUserLog(*Models.UserLog) error
}
