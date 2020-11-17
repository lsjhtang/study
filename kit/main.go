package main

import (
	httptransport "github.com/go-kit/kit/transport/http"
	"kit/services"
	"net/http"
)

func main() {
	user := new(services.User)

	endpoint := services.GetEndpoint(user)

	service := httptransport.NewServer(endpoint, services.DecodeRequest, services.EncodeResponse)

	http.ListenAndServe(":8080", service)
}


