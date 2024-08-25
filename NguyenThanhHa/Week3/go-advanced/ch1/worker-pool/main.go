package main

import (
	"fmt"
	"time"
)

func worker(queue chan int, worknum int, done chan bool, ks chan bool){
	for true {
		// dung select de cho cung luc tren ca 2 channel
		select {
			// Xu ly job trong channel queue
		case k := <- queue:
			fmt.Println("doing work!", k, "worknumber", worknum)
			done <- true

		// Neu nhan duoc signal thi return
		case <- ks:
			fmt.Println("worker halted, number", worknum)
			return
		}
	}
}

func main() {
	// channel de terminate cac worker
	killsignal := make(chan bool)

	// queue of jobs
	q := make(chan int)

	// done channel lay ra ket qua cua jobs
	done := make(chan bool)

	// So luong worker trong pool
	numOfWorkers := 4
	for i:= 0; i < numOfWorkers; i++ {
		go worker(q,i,done, killsignal)
	}

	// Dua job vao queue
	numOfJobs := 17
	for j := 0; j < numOfJobs; j++{
		go func (j int)  {
			q <- j
		}(j)
	}

	// Cho nhan du ket qua 
	for c:= 0; c < numOfJobs; c ++ {
		<- done
	}

	// Don dep cac worker
	close(killsignal)
	time.Sleep(2 * time.Second)
}