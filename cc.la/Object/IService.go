package Object

type IService interface {
	Save(data interface{}) IService
	List() IService
}