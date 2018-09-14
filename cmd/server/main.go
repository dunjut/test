package main

import (
	"log"
	"net"
	"net/http"

	"github.com/dunjut/test/rpc"
	"github.com/dunjut/test/server"
)

func main() {
	arith := new(server.Arith)
	rpc.Register(arith)
	rpc.HandleHTTP()

	l, e := net.Listen("tcp", ":1234")
	if e != nil {
		log.Fatal("listen error:", e)
	}
	if e = http.Serve(l, nil); e != nil {
		log.Fatal("serve error:", e)
	}
}
