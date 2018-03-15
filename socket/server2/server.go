package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	l, err := net.Listen("tcp", ":8888")
	if err != nil {
		fmt.Println("Listen error : ", err)
		return
	}

	defer l.Close()
	fmt.Println("Listen ok")

	var i int
	for {
		time.Sleep(time.Second * 10)

		if _, err := l.Accept(); err != nil {
			fmt.Println("Accept error : ", err)
			break
		}

		i++
		fmt.Println("Accept a new connection ", i)
	}
}
