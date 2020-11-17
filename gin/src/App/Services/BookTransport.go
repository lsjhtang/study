package Services

import (
	"github.com/gin-gonic/gin"
	"topic/src/App"
)


func CreteBookListRequest() App.EncodeRequest  {
	
	return func(context *gin.Context) (interface{}, error) {
		req := &BookListRequest{}
		err := context.BindQuery(req)
		if err != nil  {
			return nil, err
		}
		return req, nil
	}
}

func CreteBookDetailRequest() App.EncodeRequest  {

	return func(context *gin.Context) (interface{}, error) {
		req := &BookDetailRequest{}
		err := context.BindUri(req)
		if err != nil  {
			return "参数错误", err
		}
		return req, nil
	}
}

func CreteBookListResponse() App.DecodeResponse  {

	return func(context *gin.Context, res interface{})  error {
		context.JSON(200, res)
		return  nil
	}
}
