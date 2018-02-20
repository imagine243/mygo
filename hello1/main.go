package main

import (
	"errors"
	"fmt"
	"math"
	"time"
)

type geometry interface {
	area() float64
	perim() float64
}

type rect struct {
	width, height float64
}

type circle struct {
	radius float64
}

func (r rect) area() float64 {
	return r.width * r.height
}

func (r rect) perim() float64 {
	return 2*r.width + 2*r.height
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func f1(arg int) (int, error) {
	if arg == 42 {
		return -1, errors.New("can't work with 42")
	}
	return arg + 3, nil
}

type argError struct {
	arg  int
	prob string
}

func (e *argError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.prob)
}

func f2(arg int) (int, error) {
	if arg == 42 {
		return -1, &argError{arg, "can't work with it"}
	}

	return arg + 3, nil
}

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, " : ", i)
	}
}

func main() {
	// fmt.Println("hello world")
	// r := rect{3, 4}
	// c := circle{5}
	// measure(r)
	// measure(c)
	//
	// for _, i := range []int{7, 42} {
	//     if r, e := f1(i); e != nil {
	//         fmt.Println("f1 failed : ", e)
	//     } else {
	//         fmt.Println("f1 worked : ", r)
	//     }
	// }
	//
	// for _, i := range []int{7, 42} {
	//     if r, e := f2(i); e != nil {
	//         fmt.Println("f2 failed : ", e)
	//     } else {
	//         fmt.Println("f2 worked : ", r)
	//     }
	// }
	//
	// _, e := f2(42)
	// if ae, ok := e.(*argError); ok {
	//     fmt.Println(ae.arg)
	//     fmt.Println(ae.prob)
	// }

	// f("direct")
	// go f("goroutine")
	// go func(msg string) {
	//     fmt.Println(msg)
	// }("going")
	//
	// fmt.Scanln()
	// fmt.Println("done")

	// messages := make(chan string)
	// go func() {
	//     messages <- "ping"
	// }()
	//
	// msg := <-messages
	// fmt.Println(msg)

	// messages := make(chan string, 2)
	// messages <- "ping"
	// messages <- "ping"
	//
	// fmt.Println(<-messages)
	// fmt.Println(<-messages)

	// done := make(chan bool, 1)
	// go worker(done)
	// <-done

	// pings := make(chan string, 1)
	// pongs := make(chan string, 1)
	// ping(pings, "passed message")
	// pong(pings, pongs)
	// fmt.Println(<-pongs)

	// c1 := make(chan string)
	// c2 := make(chan string)
	//
	// go func() {
	//     time.Sleep(1 * time.Second)
	//     c1 <- "one"
	// }()
	//
	// go func() {
	//     time.Sleep(2 * time.Second)
	//     c2 <- "two"
	// }()
	//
	// for i := 0; i < 2; i++ {
	//     select {
	//     case msg1 := <-c1:
	//         fmt.Println("received ", msg1)
	//     case msg2 := <-c2:
	//         fmt.Println("received ", msg2)
	//
	//     }
	// }

	// c1 := make(chan string, 1)
	// go func() {
	//     time.Sleep(2 * time.Second)
	//     c1 <- "result 1"
	// }()
	//
	// select {
	// case res := <-c1:
	//     fmt.Println(res)
	// case <-time.After(1 * time.Second):
	//     fmt.Println("timeout 1")
	// }
	//
	// c2 := make(chan string, 1)
	// go func() {
	//     time.Sleep(2 & time.Second)
	//     c2 <- "reuslt 2"
	// }()
	//
	// select {
	// case res := <-c2:
	//     fmt.Println(res)
	// case <-time.After(3 * time.Second):
	//     fmt.Println("timeout 2")
	// }

	// messages := make(chan string)
	// signals := make(chan bool)
	//
	// select {
	// case msg := <-messages:
	//     fmt.Println("received msg : ", msg)
	// default:
	//     fmt.Println("no msg received")
	// }
	//
	// msg := "hi"
	// select {
	// case messages <- msg:
	//     fmt.Println("send msg : ", msg)
	// default:
	//     fmt.Println("no msg send")
	// }
	//
	// go func() {
	//     messages <- "hi"
	// }()
	//
	// for {
	//     select {
	//     case msg := <-messages:
	//         fmt.Println("received msg : ", msg)
	//     case sig := <-signals:
	//         fmt.Println("received sig : ", sig)
	//     default:
	//         fmt.Println("no msg received")
	//
	//     }
	// }

	// jobs := make(chan int, 5)
	// done := make(chan bool, 1)
	//
	// go func() {
	//     for {
	//         j, more := <-jobs
	//         if more {
	//             fmt.Println("received job : ", j)
	//         } else {
	//             fmt.Println("received all job")
	//             done <- true
	//             return
	//         }
	//     }
	// }()
	//
	// for i := 0; i < 3; i++ {
	//     jobs <- i
	//     fmt.Println("send job : ", i)
	// }
	//
	// close(jobs)
	// fmt.Println("close jobs")
	//
	// <-done

	// queue := make(chan string, 2)
	//
	// queue <- "one"
	// queue <- "two"
	// close(queue)
	//
	// for elem := range queue {
	//     fmt.Println("item : ", elem)
	// }
}

func ping(pings chan<- string, msg string) {
	pings <- msg
}

func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}

func worker(done chan bool) {
	fmt.Print("working...")
	time.Sleep(time.Second)
	fmt.Println("done")
	done <- true
}
