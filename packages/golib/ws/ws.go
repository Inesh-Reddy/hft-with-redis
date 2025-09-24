package ws

import (
	"fmt"

	"github.com/gorilla/websocket"
)


func ConnectToWs(WsUrl string) *websocket.Conn{
	conn,_,err:= websocket.DefaultDialer.Dial(WsUrl, nil)
	if err != nil{
		fmt.Println(`Error while connecting to Ws:`,err)
	}
	return conn;
}