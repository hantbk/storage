package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
)

func doClientWork(clientChan <- chan *rpc.Client){
	// Nhan vao doi tuong RPC client tu channel
	client := <- clientChan

	// Dong ket noi voi client truoc khi ham exit 
	defer client.Close()

	var reply string

	// Thuc hien loi goi RPC 
	err := client.Call("HelloService.Hello", "hello", &reply)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(reply)
}

func main() {
	// Listen tren cong 1234
	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("ListenTCP error: ", err)
	}

	clientChan := make(chan *rpc.Client)

	// Phuc vu nhieu client
	go func() {
		for {
			conn, err := listener.Accept()
			if err != nil {
				log.Fatal("Accept error: ", err)
			}
			// Khi co ket noi moi, tao 1 client moi
			// va gui toi channel clientChan
			// de doClientWork nhan duoc
			clientChan <- rpc.NewClient(conn)
		}
	}()

	doClientWork(clientChan)
}