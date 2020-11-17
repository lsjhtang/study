package main

import (
	"fmt"
	"net"
)

func main() {
	cli,err := net.Dial("tcp","127.0.0.1:8080")

	if err != nil {
		println(err.Error())
		return
	}
	l,_ :=cli.Write([]byte("1213"))
	fmt.Println(l)
}

