package models

type HandlerFunc func(u *Users)
type HandlerFuncs []HandlerFunc

func (fs HandlerFuncs) apply(u *Users) {
	for _, f := range fs {
		f(u)
	}
}
func WithID(id int) HandlerFunc {
	return func(u *Users) {
		u.ID = id
	}
}

func WithName(name string) HandlerFunc {
	return func(u *Users) {
		u.Name = name
	}
}
