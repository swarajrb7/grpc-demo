package main

import (
	"log"
	"net"
	pb "github.com/swarajrb7/grpc-demo/proto"
	"google.golang.org/grpc"
)

type helloServer struct{
	pb.GreetServiceServer
}
const(
	port = ":8080"
)

func main(){
	lis, err := net.Listen("tcp",port)
	if err != nil{
		log.Fatal("Failed to start the server %v", err)
	}
	grcpServer := grpc.NewServer()
	pb.RegisterGreetServiceServer(grcpServer, &helloServer{})
	log.Printf("Server started at %v",lis.Addr())
	if err := grcpServer.Serve(lis); err != nil{
		log.Fatal("Failed to start %v",err)
	}
}
