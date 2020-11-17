package App

import (
	"context"
	"github.com/gin-gonic/gin"
)

type Endpoint func(cxt context.Context, request interface{}) (response interface{}, err error)
type EncodeRequest func(*gin.Context) (interface{}, error)
type DecodeResponse func(*gin.Context, interface{}) error

func RegisterHandler(endpoint Endpoint, encodeFunc EncodeRequest, decodeFunc DecodeResponse) func(context *gin.Context) {

	return func(context *gin.Context) {
		req,err := encodeFunc(context)
		if err != nil {
			context.JSON(400,err.Error())
			return
		}
		req,err =endpoint(context, req)
		if err != nil {
			context.JSON(400,err.Error())
			return
		}
		err = decodeFunc(context, req)
		if err != nil {
			context.JSON(400,err.Error())
			return
		}
	}
}