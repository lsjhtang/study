package main

import (
	"cc.la/Object"
	"fmt"
	"log"
	"reflect"
	"sort"
)

type Server struct{
	Weight int
}

func (p ServerSlice) Len() int           { return len(p) }
func (p ServerSlice) Less(i, j int) bool { return p[i].Weight < p[j].Weight }
func (p ServerSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }


type ServerSlice []Server

func main() {
	s := ServerSlice{{1},{8},{2}}
	sort.Sort(s)
	log.Print(s,*&s)
	return
	/*bt := Tree.NewBTree(12)
	bt.ConnectLeft(11).ConnectRight(13).String()
	bt.Left.ConnectLeft(10).ConnectRight(9).String()
	bt.Right.ConnectLeft(8).ConnectRight(7).String()
	bt.Right.Right.ConnectRight(6).String()*/
	/*suer := []Map.User{}
	u := Map.NewUser()
	u.With("id",7).With("name","zhangshan").With("sex","男").With("phone",17620152699)
	u1 := Map.NewUser()
	u1.With("id",12).With("name","zhangshan1").With("sex","男1").With("phone",17620152699)
	u2 := Map.NewUser()
	u2.With("id",11).With("name","zhangshan2").With("sex","男2").With("phone",17620152699)

	suer = append(suer,u,u1,u2)
	sort.Slice(suer, func(i, j int) bool {
		id1 := suer[i]["id"].(int)
		id2 := suer[j]["id"].(int)
		return id1 > id2
	})
	fmt.Print(suer)*/
	/*p := Object.NewProduct(
		Object.ProductWithId(1),
		Object.ProductWithName("书本"),
		Object.ProductWithPrice(10),
		)

	var service Object.IService = Object.NewUserService()
	service.Save(p)*/

	p := new(Object.Product)
	/*p := new(Object.Product)

	r := reflect.ValueOf(p)
	if r.Kind() == reflect.Ptr {
		r = r.Elem()
	}

	values := []interface{}{1, 999, "手机"}
	
	for i:=0;i<r.NumField();i++ {
		if r.Field(i).Kind() == reflect.ValueOf(values[i]).Kind() {
			r.Field(i).Set(reflect.ValueOf(values[i]))
		}
	}*/
	m := map[string]interface{}{
		"pixel": "分辨率",
		"uid": 2,
		"Price" : 999,
		"Name": "手机",
	}
	MapToStruct(m, p)
	fmt.Print(p)

}

func MapToStruct(m map[string]interface{}, p interface{})  {
	r := reflect.ValueOf(p)
	if r.Kind() == reflect.Ptr {
		r = r.Elem()
		if r.Kind() != reflect.Struct {
			panic("must struct")
		}
		FindProperty := func(key string, tag string) interface{} {
			for k,v := range m {
				if k == key || k == tag{
					return v
				}
			}
			return nil
		}
		for i:=0;i<r.NumField();i++ {
			getValue := FindProperty(r.Type().Field(i).Name, r.Type().Field(i).Tag.Get("name"))
			if getValue != nil && reflect.ValueOf(getValue).Kind() == r.Field(i).Kind() {
				r.Field(i).Set(reflect.ValueOf(getValue))
			}
		}

	}else {
		panic(" must ptr")
	}

}
