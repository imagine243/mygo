package main

import (
	"fmt"
	"time"
)

func main() {
	// fmt.Println("hello main")
	//
	// timer1 := time.NewTimer(2 * time.Second)
	// <-timer1.C
	// fmt.Println("time 1 expired")
	//
	// timer2 := time.NewTimer(time.Second)
	// go func() {
	//     <-timer2.C
	//     fmt.Println("timer2 expired")
	// }()
	//
	// stoped := timer2.Stop()
	// if stoped {
	//     fmt.Println("timer2 stoped")
	// }

	// ticker := time.NewTicker(10 * time.Millisecond)
	// go func() {
	//     for t := range ticker.C {
	//         fmt.Println("tick at ", t)
	//     }
	// }()
	// time.Sleep(1600 * time.Millisecond)
	// ticker.Stop()
	// fmt.Println("ticker stop")

	// jobs := make(chan int, 100)
	// results := make(chan int, 100)
	//
	// for w := 1; w <= 3; w++ {
	//     go worker(w, jobs, results)
	// }
	//
	// for j := 1; j <= 5; j++ {
	//     jobs <- j
	// }
	// close(jobs)
	//
	// for a := 1; a <= 5; a++ {
	//     fmt.Println("results : ", <-results)
	// }

	// request := make(chan int, 5)
	// for i := 0; i < 5; i++ {
	//     request <- i
	// }
	// close(request)
	// limiter := time.Tick(200 * time.Millisecond)
	//
	// for req := range request {
	//     <-limiter
	//     fmt.Println("request ; ", req, time.Now())
	// }

	burstyLimiter := make(chan time.Time, 3)
	go func() {
		for i := 0; i < 3; i++ {
			burstyLimiter <- time.Now()
		}
		for t := range time.Tick(2000 * time.Millisecond) {
			burstyLimiter <- t
		}
	}()

	burstyRequests := make(chan int, 5)
	for i := 0; i < 5; i++ {
		burstyRequests <- i
	}
	// close(burstyRequests)

	go func() {
		for i := 5; i < 15; i++ {
			burstyRequests <- i
			time.Sleep(100 * time.Millisecond)
			fmt.Println("write request : ", i, time.Now())
		}
	}()

	// for req := range burstyRequests {
	//     <-burstyLimiter
	//     fmt.Println("request : ", req, time.Now())
	// }

	for {
		<-burstyLimiter
		select {
		case req := <-burstyRequests:
			fmt.Println("request : ", req, time.Now())
		default:
			fmt.Println("request over")
			return
		}
	}
}

func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("worker", id, "start job", j)
		time.Sleep(time.Second)
		fmt.Println("worker", id, "finished job", j)
		results <- j * 2
	}
}
