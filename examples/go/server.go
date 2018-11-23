package main

import (
	"flag"
	"github.com/gorilla/websocket"
	"github.com/liukaijv/ws-rpc"
	"github.com/liukaijv/ws-rpc/examples/go/data"
	"log"
	"net/http"
)

var serveAddr = flag.String("addr", "0.0.0.0:8080", "http service address")

type Chat struct {
}

func (c *Chat) Message(msg *data.Incoming, reply *data.Outputting, ws *websocket.Conn) error {
	log.Printf("recv from %v:%#v\n", ws.RemoteAddr(), msg)
	reply.Message = msg.Message
	return nil
}

func main() {

	flag.Parse()
	log.SetFlags(log.Llongfile)

	chat := Chat{}

	registry := ws_rpc.NewRegistry()

	registry.RegisterService(&chat)

	upgrader := websocket.Upgrader{}

	serve := func(w http.ResponseWriter, r *http.Request) {

		conn, err := upgrader.Upgrade(w, r, nil)

		if err != nil {
			log.Print("upgrade:", err)
			return
		}

		endpoint := ws_rpc.NewServer(conn, registry)

		go endpoint.Serve()

	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("test!"))
	})

	http.HandleFunc("/ws", serve)

	log.Printf("serve at %s", *serveAddr)

	err := http.ListenAndServe(*serveAddr, nil)
	if err != nil {
		log.Fatal(err)
	}

}
