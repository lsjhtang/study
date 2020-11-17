package Map

import (
	"fmt"
	"sort"
)

type User map[string]interface{}

func NewUser() User {
	return make(User)
}

func (this User) With(k string, v interface{}) User {
	this[k] = v
	return this
}

func (this User) Fields() []string {
	var keys []string
	for k, _ := range this {
		keys = append(keys, k)
	}
	sort.Sort(sort.StringSlice(keys))//(强制转换)
	return keys
}

func (this User) String() string  {
	str := ""
	for index,v := range this.Fields(){
		str += fmt.Sprintf("%d. %v->%v\n",index,v,this[v])
	}
	return str
}