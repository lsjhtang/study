package dao

import (
	Models "ddd/domain/models"
	"github.com/jinzhu/gorm"
)

type User struct {
	DB *gorm.DB
}

func newUser() *User {
	return &User{}
}

func (u *User) FindByName(name string) *Models.User {
	um := &Models.User{}
	u.DB.Where("user_name = ?", name).Find(um)
	return um
}

func (u *User) SaveUser(*Models.User) error {
	return nil
}

func (u *User) UpdateUser(*Models.User) error {
	return nil
}

func (u *User) DeleteUser(*Models.User) error {
	return nil
}
