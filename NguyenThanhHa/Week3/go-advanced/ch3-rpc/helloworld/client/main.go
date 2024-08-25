package main

import (
	"fmt"
	"log"
	"net/rpc"
)

func main() {
	// Ket noi den rpc server
	client, err := rpc.Dial("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("Dialing:", err)
	}
	var reply string
	// Goi RPC voi ten service da regsister
	err = client.Call("HelloService.Hello", "World", &reply)
	if err != nil {
		log.Fatal("Call HelloService.Hello error:", err)
	}
	fmt.Println(reply)
}