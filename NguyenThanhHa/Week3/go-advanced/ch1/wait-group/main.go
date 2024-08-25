package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	// Mo ra N goroutine
	for i := 0; i < 10; i++ {
		// Tang so luong su kien cho, ham nay phai duoc
		// dam bao thuc thi truoc khi bat dau 1 goroutine chay nen
		wg.Add(1)

		go func(){
			fmt.Println("Hello World")

			// Giam so luong su kien cho
			// Cho biet goroutine nay da hoan thanh
			wg.Done()
		}()

		// Doi N goroutine hoan thanh
		wg.Wait()
	}

}