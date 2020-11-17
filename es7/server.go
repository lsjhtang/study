package main

import (
	"es/Funcs"
	"github.com/gin-gonic/gin"
)

func main()  {
	r := gin.Default()
	g := r.Group("/books")
	{
		g.Handle("GET","", Funcs.BookList)
		g.Handle("GET","press/:press", Funcs.BookQuery)
		g.Handle("POST","search", Funcs.BookSearch)
	}

	helper := r.Group("/helper")
	{
		helper.Handle("GET","press", Funcs.PressList)
	}
	r.Run()
}


