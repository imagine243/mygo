package main

import "fmt"

type TTT struct{}

type person struct {
	name string
	age  int
}

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
	// twoD := make([][]int, 3):
	// for i := 0; i < 3; i++ {
	//     innerlen := i + 1
	//     twoD[i] = make([]int, innerlen)
	//     for j := 0; j < innerlen; j++ {
	//         twoD[i][j] = i + j
	//     }
	// }
	//
	// fmt.Println("2d : ", twoD)

	// m := make(map[string]int)
	// m["k1"] = 7
	// m["k2"] = 8
	//
	// fmt.Println("map : ", m)
	//
	// v1 := m["k1"]
	// fmt.Println("v1 : ", v1)
	// fmt.Println("len : ", len(m))
	//
	// delete(m, "k2")
	// fmt.Println("map : ", m)
	//
	// ff, prs := m["k2"]
	// fmt.Println("prs : ", ff, prs)
	//
	// n := map[string]int{"foo": 1, "bar": 2}
	// fmt.Println("map : ", n)

	// nums := []int{2, 3, 4}
	// sum := 0
	// for i, num := range nums {
	//     sum += num
	//     fmt.Println("index : ", i, " num : ", num)
	// }
	//
	// kvs := map[string]string{"a": "apple", "b": "banana"}
	// for k, v := range kvs {
	//     fmt.Printf("%s -> %s\n", k, v)
	// }
	//
	// for k := range kvs {
	//     fmt.Println("key : ", k)
	// }
	//
	// for i, c := range "go" {
	//     fmt.Println(i, c)
	// }

	// res := plus(1, 2)
	// fmt.Println("res : ", res)
	// res = plusPlus(1, 2, 3)
	// fmt.Println("res : ", res)
	//
	// a, b := vals()
	// fmt.Println(a)
	// fmt.Println(b)
	//
	// _, c := vals()
	// fmt.Println(c)
	//
	// total := sum(1, 2)
	// fmt.Println("total : ", total)
	// total = sum(1, 2, 3)
	// fmt.Println("total : ", total)
	// nums := []int{1, 2, 3, 4}
	// total = sum(nums...)
	// fmt.Println("total : ", total)
	//
	// nextInt := intSeq()
	//
	// fmt.Println(nextInt())
	// fmt.Println(nextInt())
	// fmt.Println(nextInt())
	//
	// newInts := intSeq()
	// fmt.Println(newInts())
	//
	// fmt.Println(fact(10))

	// i := 1
	// fmt.Println("initial : ", i)
	//
	// zeroval(i)
	// fmt.Println("zeroval : ", i)
	//
	// zeroptr(&i)
	// fmt.Println("zeroptr : ", i)
	//
	// fmt.Println("pointer : ", &i)

	// fmt.Println(person{"bob", 20})
	// fmt.Println(person{name: "alice", age: 30})
	// fmt.Println(person{name: "fred"})
	// fmt.Println(&person{name: "ann", age: 40})
	// s := person{name: "sean", age: 50}
	// fmt.Println(s.name)
	// sp := &s
	// fmt.Println(s.age)
	// sp.age = 51
	// fmt.Println(s.age)

	// r := rect{width: 10, height: 20}
	//
	// fmt.Println("area : ", r.area())
	// fmt.Println("perim : ", r.perim())
	//
	// rp := &r
	// fmt.Println("area : ", rp.area())
	// fmt.Println("perim : ", rp.perim())

}

type rect struct {
	width, height int
}

func (r *rect) area() int {
	return r.width * r.height
}

func (r rect) perim() int {
	return 2*r.width + 2*r.height
}

func zeroval(ival int) {
	ival = 0
}

func zeroptr(iptr *int) {
	*iptr = 0
}

func plus(a int, b int) int {
	return a + b
}

func plusPlus(a, b, c int) int {
	return a + b + c
}

func vals() (int, int) {
	return 3, 7
}

func sum(nums ...int) int {
	fmt.Println(nums, " ")
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
	return total
}

func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}
