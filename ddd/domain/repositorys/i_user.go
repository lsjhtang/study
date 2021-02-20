package repositorys

import (
	Models "ddd/domain/models"
)

type IUser interface {
	FindByName(name string) *Models.User
	SaveUser(*Models.User) error
	UpdateUser(*Models.User) error
	DeleteUser(*Models.User) error
}
