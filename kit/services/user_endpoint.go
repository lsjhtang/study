package services

import (
	"context"
	"github.com/go-kit/kit/endpoint"
)

type Request struct {
	Uid int `json:"uid"`
}

type Response struct {
	Status int `json:"status"`
	Msg string `json:"msg"`
	Data map[string]interface{}
}

func GetEndpoint(userServices IUserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error){
		re := request.(Request)
		result := userServices.GetName(re.Uid)
		return  Response{Data: map[string]interface{}{"result":result}}, nil
	}
}
