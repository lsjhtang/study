package util

import (
	"fmt"
	"github.com/hashicorp/consul/api"
	"strconv"
	"time"
)

var (
	apiClient   *api.Client
	err         error
	ServiceId   string
	ServiceName string
	ServicePort int
)

func init() {
	config := api.DefaultConfig()
	config.Address = "192.168.87.129:8500"
	apiClient, err = api.NewClient(config)
	if err != nil {
		panic(err)
	}
	ServiceId = "userService" + strconv.FormatInt(time.Now().UnixNano(), 10)
}

func SetNameAndPort(name string, port int) {
	ServiceName = name
	ServicePort = port
}

func RegisterService() {
	reg := api.AgentServiceRegistration{}
	reg.ID = ServiceId
	reg.Name = ServiceName
	reg.Address = "172.20.10.13"
	reg.Port = ServicePort
	reg.Tags = []string{"primary"}

	check := api.AgentServiceCheck{}
	check.Interval = "5s"
	check.HTTP = fmt.Sprintf("http://%s:%d/user/101", reg.Address, reg.Port)

	reg.Check = &check

	err = apiClient.Agent().ServiceRegister(&reg) //注册
	if err != nil {
		panic(err)
	}
}

func DeregisterService() {
	err = apiClient.Agent().ServiceDeregister(ServiceId) //反注册
	if err != nil {
		panic(err)
	}
}
