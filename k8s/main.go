package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"k8s/ginskill/common"
	"k8s/ginskill/data/mapper"
	"k8s/ginskill/models"
	"k8s/ginskill/result"
)

func main() {
	fmt.Println(mapper.NewUserMapper().GetUserList().Sql)
	return
	r := gin.Default()
	r.Use(common.ErrorHandler())
	r.GET("/", func(c *gin.Context) {
		u := models.New()
		result.Result(c.ShouldBindJSON(u)).Unwrap()
		result.OKJson(c, u)
	})
	r.Run()
}
