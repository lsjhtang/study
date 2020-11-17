package Funcs

import (
	"es/AppInit"
	"es/Models"
	"github.com/gin-gonic/gin"
	"github.com/olivere/elastic/v7"
)


//搜索
func BookSearch(ctx *gin.Context)  {
	bookSearch := models.NewSearch()
	err := ctx.BindJSON(bookSearch)
	if err != nil {
		ctx.JSON(400,gin.H{"error":err.Error()})
		return
	}

	qlist := make([]elastic.Query,0)
	if bookSearch.BookName != "" {
		mq := elastic.NewMatchQuery("BookName",bookSearch.BookName)
		qlist = append(qlist, mq)
	}

	if bookSearch.BookPress != "" {
		mq := elastic.NewTermQuery("BookPress",bookSearch.BookPress)
		qlist = append(qlist, mq)
	}

	if bookSearch.BookPrice1 > 0 && bookSearch.BookPrice2 > 0{
		mq := elastic.NewRangeQuery("BookPrice1").Gte(bookSearch.BookPrice1).Lte(bookSearch.BookPrice2)
		qlist = append(qlist, mq)
	}

	qsort := make([]elastic.Sorter,0)
	if bookSearch.OrderBy.Score {
		qsort = append(qsort, elastic.NewScoreSort().Asc())
	}

	if bookSearch.OrderBy.PriceOrder == models.ORDERBY_SCORE_ASC {
		qsort = append(qsort, elastic.NewFieldSort("BookPrice1").Asc())
	}

	if bookSearch.OrderBy.PriceOrder == models.ORDERBY_SCORE_DESC {
		qsort = append(qsort, elastic.NewFieldSort("BookPrice1").Desc())
	}

	boolquery := elastic.NewBoolQuery().Must(qlist...)
	cx,err := AppInit.GetEsClient().Search().Query(boolquery).Index("books").
		Aggregation("avg_BookPrice1",elastic.NewAvgAggregation().Field("BookPrice1")).
		SortBy(qsort...).From(0).Size(50).Do(ctx)
	if err != nil {
		ctx.JSON(400,err)
		return
	}else {
		ctx.JSON(200,gin.H{"result":cx})
	}
}
