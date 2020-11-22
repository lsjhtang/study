package util

import (
	"github.com/hashicorp/consul/api"
)

var (
	client *api.Client
	err    error
)

func init() {
	config := api.DefaultConfig()
	config.Address = "192.168.87.129:8500"
	client, err = api.NewClient(config)
	if err != nil {
		panic(err)
	}
}

func RegisterService() {
	reg := api.AgentServiceRegistration{}
	reg.ID = "userService"
	reg.Name = "userService"
	reg.Address = "172.20.10.13"
	reg.Port = 8080
	reg.Tags = []string{"primary"}

	check := api.AgentServiceCheck{}
	check.Interval = "5s"
	check.HTTP = "http://172.20.10.13:8080/user/101"

	reg.Check = &check

	err = client.Agent().ServiceRegister(&reg) //注册
	if err != nil {
		panic(err)
	}
}

func DeregisterService() {
	err = client.Agent().ServiceDeregister("userService") //反注册
	if err != nil {
		panic(err)
	}
}
