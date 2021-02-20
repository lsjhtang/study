package result

import (
	"errors"
	"github.com/gin-gonic/gin"
	"sync"
)

var ResultPool *sync.Pool

func init() {
	ResultPool = &sync.Pool{
		New: func() interface{} {
			return NewResultJson()
		},
	}
}

type ErrorResult struct {
	data interface{}
	err  error
}

func Result(es ...interface{}) *ErrorResult {
	if len(es) == 1 {
		if es[0] == nil {
			return &ErrorResult{nil, nil}
		}

		if e, ok := es[0].(error); ok {
			return &ErrorResult{nil, e}
		}
	}

	if len(es) == 2 {
		if es[1] == nil {
			return &ErrorResult{es[0], nil}
		}

		if e, ok := es[1].(error); ok {
			return &ErrorResult{es[0], e}
		}
	}
	return &ErrorResult{nil, errors.New("result error")}
}

func (e *ErrorResult) Unwrap() interface{} {
	if e.err != nil {
		panic(e.err.Error())
	}
	return e.data
}

type ResultJson struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func NewResultJson() *ResultJson {
	return &ResultJson{}
}

type OutPut func(c *gin.Context, v interface{})

func R(c *gin.Context, pams ...interface{}) func(OutPut) {
	r := ResultPool.Get().(*ResultJson)
	defer ResultPool.Put(r)

	return func(f OutPut) {
		f(c, pams)
	}
}

func OKJson(c *gin.Context, data interface{}) {
	r := ResultPool.Get().(*ResultJson)
	defer ResultPool.Put(r)
	r.Code = 200
	r.Msg = "ok"
	r.Data = data
	c.JSON(200, r)
}

func ErrorJson(c *gin.Context, msg string) {
	r := ResultPool.Get().(*ResultJson)
	defer ResultPool.Put(r)
	r.Code = 400
	r.Msg = msg
	r.Data = []interface{}{}
	c.JSON(200, r)
}

func Json(c *gin.Context, code int, msg string, data interface{}) {
	r := ResultPool.Get().(*ResultJson)
	defer ResultPool.Put(r)
	r.Code = code
	r.Msg = msg
	r.Data = data
	c.JSON(200, r)
}
