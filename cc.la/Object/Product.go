package Object

type Product struct {
	Id int `name:"uid" bbb:"qqq"`
	Price int
	Name string
}

func NewProduct(fs ...ProductAttrFunc) *Product {
	p := new(Product)
	ProductAttrFuncs(fs).apply(p)//强制转换类型fs
	return p
}