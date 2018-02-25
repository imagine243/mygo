package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("hello sort")
	// strs := []string{"c", "a", "b"}
	// sort.Strings(strs)
	// fmt.Println("strs : ", strs)
	//
	// ints := []int{8, 3, 1}
	// sort.Ints(ints)
	// fmt.Println("ints : ", ints)
	//
	// s := sort.IntsAreSorted(ints)
	// fmt.Println("Sorted : ", s)

	// fruits := []string{"peach", "banana", "kiwi"}
	// sort.Sort(byLength(fruits))
	// fmt.Println(fruits)

	panic("a problem")

	_, err := os.Create("/temp/file")
	if err != nil {
		panic(err)
	}
}

type byLength []string

func (s byLength) Len() int {
	return len(s)
}

func (s byLength) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s byLength) Less(i, j int) bool {
	return len(s[i]) < len(s[j])
}
