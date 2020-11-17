package services

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
)

func EncodeRequest(ctx context.Context, r http.Request) (interface{}, error) {

	if uid := r.URL.Query().Get("uid"); uid != ""{
		uid, _ := strconv.Atoi(uid)
		return Request{Uid: uid}, nil
	}
	return nil, errors.New("参数不存在")
}

func DecodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error  {
	w.Header().Set("Content-Type","json")
	return json.NewEncoder(w).Encode(response)

}