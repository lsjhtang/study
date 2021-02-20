package common

import (
	"github.com/gin-gonic/gin"
	"k8s/ginskill/result"
)

func ErrorHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				//ctx.JSON(400,gin.H{"msg":err})
				result.ErrorJson(ctx, err.(string))
			}
		}()

		ctx.Next()
	}
}
