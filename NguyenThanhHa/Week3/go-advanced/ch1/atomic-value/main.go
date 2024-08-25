package main

import (
	"time"
	"sync/atomic"
)

func loadConfig() map[string] string {
	return make(map[string]string)
}

func requests() chan int {
	return make(chan int)
}

func main() {

	// Nam giu thong tin config cua server
	var config atomic.Value
	// Khoi tao gia tri ban dau
	config.Store(loadConfig())

	go func ()  {
		// Cap nhat thong tin sau moi 10 giay
		for {
			time.Sleep(10 * time.Second)
			config.Store(loadConfig())
		}
	}()

	// Tao nhieu worker de xu ly request
	// Dung thong tin cau hinh gan nhat
	for i := 0; i < 10; i++ {
		go func ()  {
			for r := range requests() {
				c := config.Load()
				// Xu ly request voi thong tin cau hinh c
				_, _ = r, c
			}

		}()
	}
}