package services

import (
	"context"
	"encoding/json"
	"errors"
	mymux "github.com/gorilla/mux"
	"kit/util"
	"net/http"
	"strconv"
)

func DecodeRequest(_ context.Context, r *http.Request) (interface{}, error) {
	vars := mymux.Vars(r)
	if uid, ok := vars["uid"]; ok {
		uid, _ := strconv.Atoi(uid)
		return Request{Uid: uid}, nil
	}
	return nil, errors.New("参数不存在")
}

func EncodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(response)

}

func MyErrorEncode(ctx context.Context, err error, w http.ResponseWriter) {
	contentType, body := "text/plain; charset=utf-8", []byte(err.Error())
	w.Header().Set("Content-type", contentType)
	if e, ok := err.(*util.MyError); ok {
		w.WriteHeader(e.Code)
		_, _ = w.Write(body)
	} else {
		w.WriteHeader(500)
		_, _ = w.Write(body)
	}
}
