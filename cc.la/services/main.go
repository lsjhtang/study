package main

import (
	"fmt"
	"net"
)

func main() {

	lists,err := net.Listen("tcp","127.0.0.1:8080")

	if err != nil {
		println(err.Error())
		return
	}
	defer lists.Close()
	client,err := lists.Accept()
	if err != nil {
		println(err.Error())
		return
	}
	defer client.Close()
	buf := make([]byte, 512)
	n,err := client.Read(buf)
	if err != nil {
		println(err.Error())
		return
	}
	fmt.Printf("%d,%s", n, buf)

}

