package main

import (
	"flag"
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

	endpoint, err := ws_rpc.NewClient(u.String(), nil)

	if err != nil {
		log.Fatal("NewClient:", err)
	}

	defer endpoint.Close()

	var reply data.Outputting

	err = endpoint.Call("Chat.Message", &data.Incoming{From: "Tom", Message: "hello!"}, &reply)
	if err != nil {
		log.Fatal("Call:", err)
	}

	log.Print("recv msg:", reply.Message)

}
