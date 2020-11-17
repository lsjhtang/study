package Funcs

import (
	"es/AppInit"
	"github.com/gin-gonic/gin"
	"github.com/olivere/elastic/v7"
)



func MapToSilce(re *elastic.SearchResult, key string) []interface{} {
	results := make([]interface{},0)
	for _, hit := range re.Hits.Hits {
		results=append(results, hit.Fields[key].([]interface{})[0])
	}
	return  results
}

func PressList(ctx *gin.Context) {
	bp := elastic.NewCollapseBuilder("BookPress")
	cx,err := AppInit.GetEsClient().Search().Collapse(bp).FetchSource(false).Index("books").Size(50).Do(ctx)
	if err != nil {
		ctx.JSON(400,err)
		return
	}else {
		ctx.JSON(200,gin.H{"result": MapToSilce(cx,"BookPress")})
	}
}

