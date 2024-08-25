package main

import (
	"fmt"
	"log"
	"net/rpc"
	"time"
)

func doClientWork(client *rpc.Client){
	// Khoi tao 1 Goroutine rieng biet de giam sat
	// khoa thay doi
	go func() {
		var keyChanged string
		// loi goi watch synchronous se block 
		// cho den khi co su thay doi hoac timeout
		err := client.Call("KVStoreService.Watch", 30, &keyChanged)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("watch:", keyChanged)
	}()

	err := client.Call(
		// Ham Set se thay doi gia tri cua key "abc"
		"KVStoreService.Set", [2]string{"abc", "value 1"}, 
		new(struct{}),
	)

	err = client.Call(
		// Ham Set se thay doi gia tri cua key "abc"
		"KVStoreService.Set", [2]string{"abc", "value 2"},
		new(struct{}),
	)
	if err != nil {
		log.Fatal(err)
	}
	time.Sleep(3 * time.Second)
}

func main() {
	client, err := rpc.Dial("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("dialing:", err)
	}

	doClientWork(client)
}