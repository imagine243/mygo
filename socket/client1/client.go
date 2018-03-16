package main

import (
	"fmt"
	"net"
)

func main() {
	fmt.Println("begin dial")
	conn, err := net.Dial("tcp", ":8888")
	if err != nil {
		fmt.Println("dial err : ", err)
		return
	}

	defer conn.Close()
	fmt.Println("dial ok")
}
