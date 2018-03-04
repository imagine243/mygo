package main

import "fmt"

func main() {
	s := make([]int, 10, 100)
	var s1 *[]int = new([]int)
	fmt.Println(s)
	fmt.Println(*s1)
}
