package main

import (
	"flag"
	"github.com/gorilla/websocket"
	"github.com/liukaijv/ws-rpc"
	"github.com/liukaijv/ws-rpc/examples/go/data"
	"log"
	"net/url"
)

var addr = flag.String("addr", "localhost:8080", "http service address")

func main() {

	flag.Parse()
	log.SetFlags(log.Llongfile)

	u := url.URL{Scheme: "ws", Host: *addr, Path: "/ws"}
	log.Printf("connecting to %s", u.String())

	conn, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer conn.Close()

	endpoint := ws_rpc.NewClient(conn)

	var reply data.Outputting

	err = endpoint.Call("Chat.Message", &data.Incoming{From: "Tom", Message: "hello!"}, &reply)
	if err != nil {
		log.Fatal("Call:", err)
	}

	log.Print("recv msg:", reply.Message)

}
