package Object

import "fmt"

type ProductService struct {

}


func (this *ProductService) Save(data interface{}) IService  {
	fmt.Print("UserService")
	return this
}

func (this *ProductService) List() IService  {
	fmt.Print("UserService")
	return this
}