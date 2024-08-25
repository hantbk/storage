package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	// "time"
)

// producer: lien tuc tao ra 1 chuoi so nguyen dua tren boi so factor va gui chung vao channel out
func Producer(factor int, out chan<- int){
	for i := 0; ; i++ {
		out <- i * factor
	}
}

// consumer: lien tuc nhan cac so nguyen tu channel ra de print
func Consumer(in <-chan int){
	for v := range in {
		fmt.Println(v)
	}
}

func main() {
	// Hang doi
	ch := make(chan int, 64)
	// consumer
	go Producer(3, ch)
	go Producer(5, ch)

	go Consumer(ch)

	// Thoat ra sau khi chay trong 1 giay
	// time.Sleep(1 * time.Second)

	sig := make(chan os.Signal, 1)

	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)

	fmt.Printf("quit (%v)\n", <-sig)
}