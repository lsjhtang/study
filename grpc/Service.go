package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"grpc/services"
	"log"
	"net"
)

/*func main()  {
	rpcService := grpc.NewServer()

	services.RegisterProdServiceServer(rpcService, new(services.ProdService))

	ls,_ :=net.Listen("tcp", ":8081")

	_ = rpcService.Serve(ls)
}
*/
func Service()  {
	tra,err := credentials.NewServerTLSFromFile("F:/phpstudy_pro/Extensions/Nginx1.15.11/conf/ssl/aa.cc.pem", "F:/phpstudy_pro/Extensions/Nginx1.15.11/conf/ssl/aa.cc.key")
	if err != nil {
		log.Fatal(err)
	}
	rpcService := grpc.NewServer(grpc.Creds(tra))

	services.RegisterProdServiceServer(rpcService, new(services.ProdService))

	ls,_ :=net.Listen("tcp", ":8081")

	_ = rpcService.Serve(ls)
}

func main()  {
	//Client()
	Service()
}

func Client()  {
	tra,err := credentials.NewClientTLSFromFile("F:/phpstudy_pro/Extensions/Nginx1.15.11/conf/ssl/aa.cc.pem", "aa.cc")
	if err != nil {
		log.Fatal(err)
	}

	client,err := grpc.Dial(":8081", grpc.WithTransportCredentials(tra))
	if err != nil {
		log.Println(err)
	}
	defer client.Close()

	prodClient := services.NewProdServiceClient(client)
	resResult,err := prodClient.GetProdStock(context.Background(), &services.Request{ProdId: 2, ProdArea: 1})
	resResults,err := prodClient.GetProdStocks(context.Background(), &services.QuerySize{Size: 2})
	if err != nil {
		log.Panic(err)
	}
	fmt.Println(resResult,resResults)
}