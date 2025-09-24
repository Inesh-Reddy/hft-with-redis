package main

import (
	"fmt"
	"log"
	"net"

	"github.com/gorilla/websocket"
	"google.golang.org/grpc"
)

func BinanceWs(WsUrl)*websocket.Conn{
	ConnectToWs(WsUrl)
}
func main(){
	fmt.Println(`Ticker servive running ....`)

	lis, err :=net.Listen("tcp",":50051");
	if err != nil{
		log.Println("error while connecting to TCP layer: ", err)
	}

	grpc := grpc.NewServer()
	log.Println("Go server listening on port: 50051....")
	if err :=grpc.Serve(lis)
	err != nil {
		fmt.Print("Failed to serve grpc",err)
	}
}
