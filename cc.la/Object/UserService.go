package Object

import "fmt"

type UserService struct {

}

func NewUserService() *UserService {
	return &UserService{}
}

func (this *UserService) Save(data interface{}) IService {
	if pro,ok := data.(*Product); ok {
		fmt.Print(*pro)
	}else {
		fmt.Print("失败")
	}
	return this
}

func (this *UserService) List() IService {
	fmt.Print("UserService")
	return this
}