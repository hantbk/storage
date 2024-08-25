package main

import "fmt"

func main() {
	// done := make(chan int, 1)

	// go func() {
	// 	fmt.Println("Hello World")

	// 	// Gui 1 gia tri vao channel done
	// 	done <- 1
	// }()

	// // main thread nhan gia tri tu channel done
	// // trang thai block
	// <-done

	done := make(chan int, 10)

	// Mo ra N goroutine
	for i := 0; i < cap(done); i++ {
		go func(){
			fmt.Println("Hello World")
			done <- 1
		}()
	}

	// Doi ca 10 goroutine hoan thanh
	for i := 0; i < cap(done); i++ {
		<-done
	}
}

