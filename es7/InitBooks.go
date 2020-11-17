package main

import (
	"context"
	"es/AppInit"
	"fmt"
	"github.com/olivere/elastic/v7"
	"strconv"
	"sync"
	"wserver/models"
)

func main()  {


	page:=1
	pagesize:=2000
	wg := sync.WaitGroup{}
	for{
		book_list:= models.BookList{}
		db:=AppInit.GetDB().Order("book_id desc").Limit(pagesize).Offset((page-1)*pagesize).Find(&book_list)
		if db.Error!=nil || len(book_list)==0{
			break
		}

		wg.Add(1)
		go func() {
			defer wg.Done()
			client := AppInit.GetEsClient()
			bulk := client.Bulk()
			for _,book:=range book_list {
				es := elastic.NewBulkIndexRequest().Index("books").Id(strconv.Itoa(book.BookID)).Doc(book)
				bulk.Add(es)
			}

			rsp,err:=bulk.Do(context.Background())
			if err!=nil{
				fmt.Println(err)
			}else {
				fmt.Println(rsp)
			}
		}()

		page++
	}

	wg.Wait()
}
