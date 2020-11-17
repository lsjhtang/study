package services

import (
	"context"
	"github.com/go-kit/kit/endpoint"
)

type Request struct {
	Uid int `json:"uid"`
}

type Response struct {
	Result string `json:"result"`
}

func GetEndpoint(userServices IUserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error){
		re := request.(Request)
		result := userServices.GetName(re.Uid)
		return  Response{Result:result}, nil
	}
}
