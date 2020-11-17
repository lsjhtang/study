package Services

import (
	"context"
	"topic/src/App"
)

type BookListRequest struct {
	Size int `form:"size"`
}

type BookResponse struct {
	Msg string `json:"msg"`
	Data interface{} `json:"data"`
}
type Date map[string]interface{}

type BookDetailRequest struct {
	BookID int `uri:"id" binding:"required,gt=0,max=100000"`
}

func BookListEndPoint(book *BookService) App.Endpoint {
	return func(cxt context.Context, request interface{}) (response interface{}, err error) {
		req := request.(*BookListRequest)
		return &BookResponse{Data:book.GetBookList(req)}, nil
	}
}

func BookDetailEndPoint(book *BookService) App.Endpoint {
	return func(cxt context.Context, request interface{}) (interface{}, error) {
		req := request.(*BookDetailRequest)
		book, bookMeta, err := book.GetBookDetail(req)
		return &BookResponse{Data:Date{"book":book, "bookMeta":bookMeta}}, err
	}
}