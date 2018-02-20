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

	jobs := make(chan int, 100)
	results := make(chan int, 100)

	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	for j := 1; j <= 5; j++ {
		jobs <- j
	}
	close(jobs)

	for a := 1; a <= 5; a++ {
		fmt.Println("results : ", <-results)
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
