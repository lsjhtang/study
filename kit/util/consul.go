package util

import (
	"github.com/hashicorp/consul/api"
)

func RegisterService() {
	config := api.DefaultConfig()
	config.Address = "192.168.87.128:8500"

	reg := api.AgentServiceRegistration{}
	reg.ID = "userService"
	reg.Name = "userService"
	reg.Address = "192.168.160.1"
	reg.Port = 8080
	reg.Tags = []string{"primary"}

	check := api.AgentServiceCheck{}
	check.Interval = "5s"
	check.HTTP = "http://192.168.160.1:8080/user/101"

	reg.Check = &check

	client, err := api.NewClient(config)
	if err != nil {
		panic(err)
	}
	err = client.Agent().ServiceRegister(&reg) //注册
	//err = client.Agent().ServiceDeregister("userService")//取消
	if err != nil {
		panic(err)
	}
}
