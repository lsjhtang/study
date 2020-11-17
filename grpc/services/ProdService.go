package services

import (
	"context"
)

type ProdService struct {

}


func(this *ProdService) GetProdStock(ctx context.Context, req *Request) (*Response, error) {
	var prod_area int32
	if req.ProdArea == ProAreas_A {
		prod_area = 44
	}else if req.ProdArea == ProAreas_B {
		prod_area = 55
	}else if req.ProdArea == ProAreas_C {
		prod_area = 66
	}
	return &Response{ProdStock: prod_area}, nil
}

func(this *ProdService) GetProdStocks(ctx context.Context, size *QuerySize) (*ResponseList, error) {
	list := []*Response{{ProdStock: 1}, {ProdStock: 2}, {ProdStock: 3}}
	return &ResponseList{Prodreq: list}, nil
}

func(this *ProdService) GetProdInfo(context.Context, *Request) (*ProdModel, error) {
		prodm := &ProdModel{
			prod_id:1,
			prod_name:"测试",
			prod_prize:3.6,
		}
	return prodm, nil
}
