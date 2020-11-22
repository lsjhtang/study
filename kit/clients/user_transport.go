package clients

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func EncodeRequest(_ context.Context, request *http.Request, r interface{}) error {
	req := r.(Request)
	request.URL.Path += "/user/" + strconv.Itoa(req.Uid)
	return nil
}

func DecodeResponst(_ context.Context, r *http.Response) (response interface{}, err error) {
	if r.StatusCode >= 400 {
		return nil, fmt.Errorf("%s", "数据异常")
	}
	resp := &Response{}
	err = json.NewDecoder(r.Body).Decode(resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
