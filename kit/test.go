package main

import (
	"fmt"
	"github.com/afex/hystrix-go/hystrix"
	"math/rand"
	"time"
)

func main() {
	hystrix.ConfigureCommand("my_command", hystrix.CommandConfig{
		Timeout:               2000,
		MaxConcurrentRequests: 100,
		ErrorPercentThreshold: 25,
		RequestVolumeThreshold: 3,
	})

	for  {
		output := make(chan interface{},1)
		errors := hystrix.Go("my_command", func() error {
			if rand.Intn(10) > 6 {
				output <- Products{Id: 1, Name: "go进阶成功", Price: 1000}
				time.Sleep(time.Second*1)
			}else {
				output <- Products{Id: 2, Name: "go进阶失败", Price: 0}
				time.Sleep(time.Second*2)
			}
			return nil
		}, func(err error) error {
			output <- Products{Id: 3, Name: "服务降级", Price: 0}
			return err
		})

		select {
		case result:= <-output:
			time.Sleep(time.Second*1)
			fmt.Println(result)
		case result:= <-errors:
			time.Sleep(time.Second*1)
			fmt.Println(result)
		}
	}
}

type Products struct {
	Id int
	Name string
	Price int
}
