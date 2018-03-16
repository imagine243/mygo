package main

import (
	"fmt"
	"net"
	"time"
)

func establishConn(i int) net.Conn {
	conn, err := net.DialTimeout("tcp", ":8888", 2*time.Second)
	if err != nil {
		fmt.Println("dial error ", i, err)
		return nil
	}
	fmt.Println("connect to server ok ", i)
	return conn
}

func main() {
	var s1 []net.Conn

	for i := 0; i < 1000; i++ {
		conn := establishConn(i)
		s1 = append(s1, conn)
	}
}
