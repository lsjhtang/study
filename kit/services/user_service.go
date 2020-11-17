package services

type IUserService interface {
	GetName(id int) string
}

type User struct {

}

func (u *User) GetName(id int) string {
	if id > 0 {
		return "abc"
	}
	return "string(id)"
}