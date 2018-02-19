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
}

func worker(done chan bool) {
	fmt.Print("working...")
	time.Sleep(time.Second)
	fmt.Println("done")
	done <- true
}
