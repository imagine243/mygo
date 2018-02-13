package main

import "fmt"
import "time"

type TTT struct{}

func main() {
	// fmt.Println("hello world")
	// fmt.Println("go " + "lang")
	// fmt.Println("1+1 = ", 1+1)
	// fmt.Println("7.0 / 3.0 = ", 7.0/3.0)
	// fmt.Println(true && false)
	// fmt.Println(true || false)
	// fmt.Println(!true)

	// i := 1
	// for i <= 3 {
	//     fmt.Println(i)
	//     i++
	// }
	//
	// for j := 5; j <= 9; j++ {
	//     fmt.Println(j)
	// }
	//
	// for {
	//     fmt.Println("loop")
	//     break
	// }
	//
	// for n := 0; n <= 12; n++ {
	//     if n%2 == 0 {
	//         continue
	//     }
	//     fmt.Println(n)
	// }

	i := 2
	fmt.Print("write ", i, " as ")
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("this is the weekend")
	default:
		fmt.Println("it is a weekday")
	}

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("it's before noon")
	default:
		fmt.Println("it's afert noon")
	}

	whoami := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("i am a bool")
		case int:
			fmt.Println("i am a int")
		default:
			fmt.Printf("don't konw type %T", t)
		}
	}

	whoami(true)
	whoami(5)
	whoami("hey")
}
