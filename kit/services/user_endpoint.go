package services

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"golang.org/x/time/rate"
	"kit/util"
	"net/http"
	"strconv"
)

type Request struct {
	Uid int `json:"uid"`
}

type Response struct {
	Status int    `json:"status"`
	Msg    string `json:"msg"`
	Data   map[string]interface{}
}

func GetEndpoint(userServices IUserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		re := request.(Request)
		result := userServices.GetName(re.Uid) + ":" + strconv.Itoa(util.ServicePort)
		return Response{Status: http.StatusOK, Data: map[string]interface{}{"result": result}}, nil
	}
}

func RateLimit(limit *rate.Limiter) endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {

		return func(ctx context.Context, request interface{}) (response interface{}, err error) {
			if !limit.Allow() {
				return nil, util.NewMyError(429, "错误", []interface{}{})
			}
			return next(ctx, request)
		}

	}
}
