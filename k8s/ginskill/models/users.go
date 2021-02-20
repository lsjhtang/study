package models

type Users struct {
	ID   int    `json:"id"`
	Name string `json:"name" binding:"required"`
}

func New(fs ...HandlerFunc) *Users {
	u := &Users{}
	HandlerFuncs(fs).apply(u)
	return u
}

func (u *Users) Mutate(fs ...HandlerFunc) *Users {
	HandlerFuncs(fs).apply(u)
	return u
}
