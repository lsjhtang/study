package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		var ips string
		if getIp() == "192.168.63.133" {
			ips = "main"
		} else {
			ips = "work1"
		}
		c.JSON(200, gin.H{
			"message": ips,
		})
	})
	err := r.Run(":8080") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	if err != nil {
		log.Fatal(err)
	}
}

func getIp() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return ""
	}
	for _, value := range addrs {
		if ipnet, ok := value.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}
	return ""
}
