package main

import "fmt"

type TTT struct{}

func main() {
	// fmt.Println("hello world")
	// fmt.Println("go " + "lang")
	// fmt.Println("1+1 = ", 1+1)
	// fmt.Println("7.0 / 3.0 = ", 7.0/3.0)
	// fmt.Println(true && false)
	// fmt.Println(true || false)
	// fmt.Println(!true)

	i := 1
	for i <= 3 {
		fmt.Println(i)
		i++
	}

	for j := 5; j <= 9; j++ {
		fmt.Println(j)
	}

	for {
		fmt.Println("loop")
		break
	}

	for n := 0; n <= 12; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
	ntl
}
