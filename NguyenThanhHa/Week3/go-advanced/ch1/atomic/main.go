package main

import (
	"fmt"
	"sync"
)

// atomic struct
var total struct {
	sync.Mutex
	value int
}

// Co che lock va unlock sau khi truy nhap vao doan gang
func worker(wg *sync.WaitGroup){
	// Thong bao goroutine da hoan thanh
	// khi ra khoi ham
	defer wg.Done()

	for i := 0; i <= 100; i++ {
		// Lock
		total.Lock()
		// Dam bao lenh total.value += i la atomic
		total.value += i
		// Unlock
		total.Unlock()
	}
}

func main() {
	// Khai bao wg de main Goroutine dung cho cac Goroutine khac khi ket thuc chuong trinh
	var wg sync.WaitGroup
	// wg can cho 2 goroutine khac
	wg.Add(2)
	
	go worker(&wg)

	go worker(&wg)

	wg.Wait()

	fmt.Println(total.value)
}