package Object

type ProductAttrFunc func(p *Product)
type ProductAttrFuncs []ProductAttrFunc

func (this ProductAttrFuncs) apply(p *Product)  {
	for _,f := range this{
		f(p)
	}
}

func ProductWithId(id int) ProductAttrFunc {
	return func(p *Product) {
		p.Id = id
	}
}

func ProductWithName(name string) ProductAttrFunc {
	return func(p *Product) {
		p.Name = name
	}
}

func ProductWithPrice(price int) ProductAttrFunc {
	return func(p *Product) {
		p.Price = price
	}
}
