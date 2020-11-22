package main

import (
	"context"
	"fmt"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/sd"
	"github.com/go-kit/kit/sd/consul"
	httptransport "github.com/go-kit/kit/transport/http"
	mymux "github.com/gorilla/mux"
	"github.com/hashicorp/consul/api"
	"io"
	"kit/clients"
	"kit/services"
	"kit/util"
	"net/http"
	"net/url"
	"os"
	"os/signal"
	"syscall"
)

func main1() {
	user := new(services.User)

	endpoints := services.GetEndpoint(user)

	service := httptransport.NewServer(endpoints, services.DecodeRequest, services.EncodeResponse)

	router := mymux.NewRouter()
	router.Handle(`/user/{uid:\d+}`, service)

	chann := make(chan error)
	go func() {
		util.RegisterService() //注册服务

		err := http.ListenAndServe(":8080", router)
		if err != nil {
			chann <- err
		}
	}()

	go func() {
		sig := make(chan os.Signal)
		signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
		chann <- fmt.Errorf("%s", <-sig)
	}()

	err := <-chann
	if err != nil {
		util.DeregisterService() //反注册

		panic(err)
	}
}

func main() { //去注册中心取服务地址然后调用服务api
	config := api.DefaultConfig()
	config.Address = "192.168.87.129:8500"
	apiClient, err := api.NewClient(config)
	if err != nil {
		panic(err)
	}
	client := consul.NewClient(apiClient)

	logger := log.NewLogfmtLogger(os.Stdout)

	tags := []string{"primary"}

	instancer := consul.NewInstancer(client, logger, "userService", tags, true)

	f := func(instance string) (endpoint.Endpoint, io.Closer, error) {
		tgt, _ := url.Parse("http://" + instance)
		return httptransport.NewClient("GET", tgt, clients.EncodeRequest, clients.DecodeResponst).Endpoint(), nil, nil

	}
	endpointer := sd.NewEndpointer(instancer, f, logger)
	endpoints, _ := endpointer.Endpoints()
	response := endpoints[0]
	res, err := response(context.Background(), clients.Request{Uid: 100})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(*res.(*clients.Response))
}

func maina() {
	tgt, _ := url.Parse("http://localhost:8080/")
	cli := httptransport.NewClient("GET", tgt, clients.EncodeRequest, clients.DecodeResponst)

	endpoints := cli.Endpoint()
	res, err := endpoints(context.Background(), clients.Request{Uid: 100})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(*res.(*clients.Response))
}
