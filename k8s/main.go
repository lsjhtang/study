package main

import (
	"github.com/gin-gonic/gin"
	"k8s/events/eventBus"
	"k8s/ginskill/common"
	"k8s/ginskill/data/mapper"
	"k8s/ginskill/result"
)

func getList(id int) string {
	return mapper.NewUserMapper().GetUserList(id).Sql
}

func main() {
	r := gin.Default()
	r.Use(common.ErrorHandler())
	r.GET("/", func(c *gin.Context) {
		/*u := models.New()
		result.Result(c.ShouldBindJSON(u)).Unwrap()*/
		bus := eventBus.NewEventBus()
		cs := bus.Sub("abc", getList)
		bus.Pub("abc", 154)
		data := <-cs
		result.OKJson(c, data.Data)
	})
	r.Run()
}
