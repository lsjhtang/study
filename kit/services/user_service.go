package services

type IUserService interface {
	GetName(id int) string
}

type User struct {

}

func (u *User) GetName(id int) string {
	return string(id)
}