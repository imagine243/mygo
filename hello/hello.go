package main

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

	// i := 2
	// fmt.Print("write ", i, " as ")
	// switch i {
	// case 1:
	//     fmt.Println("one")
	// case 2:
	//     fmt.Println("two")
	// case 3:
	//     fmt.Println("three")
	// }
	//
	// switch time.Now().Weekday() {
	// case time.Saturday, time.Sunday:
	//     fmt.Println("this is the weekend")
	// default:
	//     fmt.Println("it is a weekday")
	// }
	//
	// t := time.Now()
	// switch {
	// case t.Hour() < 12:
	//     fmt.Println("it's before noon")
	// default:
	//     fmt.Println("it's afert noon")
	// }
	//
	// whoami := func(i interface{}) {
	//     switch t := i.(type) {
	//     case bool:
	//         fmt.Println("i am a bool")
	//     case int:
	//         fmt.Println("i am a int")
	//     default:
	//         fmt.Printf("don't konw type %T", t)
	//     }
	// }
	//
	// whoami(true)
	// whoami(5)
	// whoami("hey")

	// var a [5]int
	// fmt.Println("emp: ", a)
	//
	// a[4] = 100
	// fmt.Println("set : ", a)
	// fmt.Println("get : ", a[4])
	// fmt.Println("len : ", len(a))
	//
	// b := [5]int{1, 2, 3, 4, 5}
	// fmt.Println("dcl : ", b)
	//
	// var twoD [2][3]int
	// for i := 0; i < 2; i++ {
	//     for j := 0; j < 3; j++ {
	//         twoD[i][j] = i + j
	//     }
	// }
	// fmt.Println("2d : ", twoD)

	// s := make([]string, 3)
	// fmt.Println("emp : ", s)
	//
	// s[0] = "a"
	// s[1] = "b"
	// s[2] = "c"
	// fmt.Println("set : ", s)
	// fmt.Println("get : ", s[2])
	//
	// fmt.Println("len : ", len(s))
	//
	// s = append(s, "d")
	// s = append(s, "e", "f")
	// fmt.Println("apd : ", s)
	//
	// c := make([]string, len(s))
	// copy(c, s)
	// fmt.Println("cpy : ", c)
	//
	// l := s[2:5]
	// fmt.Println("sl1 : ", l)
	//
	// l = s[:5]
	// fmt.Println("sl2 : ", l)
	//
	// l = s[2:]
	// fmt.Println("sl3 : ", l)
	//
	// t := []string{"g", "h", "l"}
	// fmt.Println("dcl : ", t)
	//
	// twoD := make([][]int, 3)
	// for i := 0; i < 3; i++ {
	//     innerlen := i + 1
	//     twoD[i] = make([]int, innerlen)
	//     for j := 0; j < innerlen; j++ {
	//         twoD[i][j] = i + j
	//     }
	// }
	//
	// fmt.Println("2d : ", twoD)
}
