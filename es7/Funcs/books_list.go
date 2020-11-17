package Funcs

import (
	"es/AppInit"
	"es/Models"
	"github.com/gin-gonic/gin"
	"github.com/olivere/elastic/v7"
	"reflect"
)

func ToMap(re *elastic.SearchResult) []*models.Books {
	booklist := models.BookList{}
	var book *models.Books
	for _,result :=	range re.Each(reflect.TypeOf(book)) {
		booklist = append(booklist, result.(*models.Books))
	}
	return  booklist
}

func BookList(ctx *gin.Context)  {
	cx,err := AppInit.GetEsClient().Search().Index("books").Do(ctx)
	if err != nil {
		ctx.JSON(400,err)
		return
	}else {
		ctx.JSON(200,gin.H{"result": ToMap(cx)})
	}
}

func BookQuery(ctx *gin.Context)  {
	param,_ := ctx.Params.Get("press")
	query := elastic.NewTermQuery("BookPress", param)
	cx,err := AppInit.GetEsClient().Search().Query(query).Index("books").Do(ctx)
	if err != nil {
		ctx.JSON(400,err)
		return
	}else {
		ctx.JSON(200,gin.H{"result": ToMap(cx)})
	}
}