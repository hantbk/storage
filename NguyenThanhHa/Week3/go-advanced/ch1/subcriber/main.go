package main

import (
	"example/hant/pubsub"
	"time"
	"strings"
	"fmt"
)

func main() {
	// Khoi tao 1 publisher
	p := pubsub.NewPublisher(100 * time.Millisecond, 10)

	// Dam bao p duoc dong truoc khi exit
	defer p.Close()

	// all subcriber het tat ca topic
	all := p.Subscribe()

	// Subcribe cac topic co "golang"
	golang := p.SubscribeTopic(func(v interface{}) bool {
		s, ok := v.(string)
		if !ok {
			return false
		}
		return strings.Contains(s, "golang")
	})

	// publish 2 topic
	p.Publish("hello, world!")
	p.Publish("hello, golang")

	// Print nhung gi subcriber all nhan duoc
	go func(){
		for msg := range all {
			fmt.Println("all: ", msg)
		}
	}()

	// Print nhung gi subcriber golang nhan duoc
	// Print nhung gi subcriber all nhan duoc
	go func(){
		for msg := range golang {
			fmt.Println("golang: ", msg)
		}
	}()

	// Thoat ra sau 3 s
	time.Sleep(3 * time.Second)
}
