package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/sd"
	"github.com/go-kit/kit/sd/consul"
	"github.com/go-kit/kit/sd/lb"
	httptransport "github.com/go-kit/kit/transport/http"
	mymux "github.com/gorilla/mux"
	"github.com/hashicorp/consul/api"
	"golang.org/x/time/rate"
	"io"
	"kit/clients"
	"kit/services"
	"kit/util"
	"net/http"
	"net/url"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"
)

func main3() {
	name := flag.String("name", "", "服务名称")
	port := flag.Int("p", 0, "服务端口")
	flag.Parse()
	if *name == "" {
		panic("请指定服务名")
	}
	if *port == 0 {
		panic("请指定端口号")
	}
	util.SetNameAndPort(*name, *port)

	user := new(services.User)

	limit := rate.NewLimiter(1, 6)
	endpoints := services.RateLimit(limit)(services.GetEndpoint(user))

	//endpoints := services.GetEndpoint(user)
	options := []httptransport.ServerOption{ //自定义错误
		httptransport.ServerErrorEncoder(services.MyErrorEncode),
	}

	service := httptransport.NewServer(endpoints, services.DecodeRequest, services.EncodeResponse, options...)

	router := mymux.NewRouter()
	router.Handle(`/user/{uid:\d+}`, service)

	chann := make(chan error)
	go func() {
		util.RegisterService() //注册服务

		err := http.ListenAndServe(":"+strconv.Itoa(*port), router)
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

		fmt.Println(err)
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

	mylb := lb.NewRoundRobin(endpointer) //内置轮询算法
	for {
		response, _ := mylb.Endpoint()
		res, err := response(context.Background(), clients.Request{Uid: 100})
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(*res.(*clients.Response))
		time.Sleep(time.Second * 3)
	}
}

func main2() {
	tgt, _ := url.Parse("http://localhost:8080/")
	cli := httptransport.NewClient("GET", tgt, clients.EncodeRequest, clients.DecodeResponst)

	endpoints := cli.Endpoint()
	res, err := endpoints(context.Background(), clients.Request{Uid: 100})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(*res.(*clients.Response))
}
