package main

import s "strings"
import "fmt"

var p = fmt.Println

func main() {
	fmt.Println("hello string")

	p("Contains", s.Contains("test", "es"))
	p("count", s.Count("test", "t"))
	p("HasPrefix", s.HasPrefix("test", "te"))
	p("HasSuffix", s.HasSuffix("test", "st"))
	p("Index", s.Index("test", "e"))
	p("join", s.Join([]string{"a", "b", "c"}, "-"))
	p("repeat", s.Repeat("a", 5))
	p("replaces", s.Replace("foo", "o", "0", -1))
	p("replaces", s.Replace("foo", "o", "0", 1))
	p("split", s.Split("a-b-c-d-e", "-"))
	p("tolower", s.ToLower("TEST"))
	p("ToUpper", s.ToUpper("test"))
	p()
	p("len", len("hello"))
	p("char", "hello"[3])
}
