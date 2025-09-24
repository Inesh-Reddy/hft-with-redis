package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/Inesh-Reddy/hft-with-redis/packages/golib/redis"
	"github.com/Inesh-Reddy/hft-with-redis/packages/golib/ws"
	"github.com/gorilla/websocket"
	"google.golang.org/grpc"
)

func BinanceWs(WsUrl string)*websocket.Conn{
	binance:=ws.ConnectToWs(WsUrl)
	return binance
}


func main(){
	fmt.Println(`Ticker servive running ....`)
	ctx := context.Background()

	lis, err :=net.Listen("tcp",":50051");
	if err != nil{
		log.Println("error while connecting to TCP layer: ", err)
	}

	binance:=BinanceWs("wss://stream.binance.com:9443/ws/btcusdt@ticker")
	_,msg,err:=binance.ReadMessage()
	if err != nil{
		fmt.Println("Erro while reading data from binanace: ",err)
	}
	fmt.Println(string(msg))

	BinanceRedis := redis.ConnectToRedis()
	BinanceRedis.Set(ctx, "ticker-active-subscribers-clientId", "btcusdt")
	
	grpc := grpc.NewServer()
	log.Println("Go server listening on port: 50051....")
	if err :=grpc.Serve(lis)
	err != nil {
		fmt.Print("Failed to serve grpc",err)
	}
}
