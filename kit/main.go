package main

import (
	httptransport "github.com/go-kit/kit/transport/http"
	mymux "github.com/gorilla/mux"
	"kit/services"
	"kit/util"
	"net/http"
)

func main() {
	user := new(services.User)

	endpoint := services.GetEndpoint(user)

	service := httptransport.NewServer(endpoint, services.DecodeRequest, services.EncodeResponse)

	router := mymux.NewRouter()
	router.Handle(`/user/{uid:\d+}`, service)

	util.RegisterService() //注册服务

	http.ListenAndServe(":8080", router)
}
