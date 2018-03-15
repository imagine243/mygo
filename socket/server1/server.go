package main

import (
	"fmt"
	"io/ioutil"
	"net"
)

func main() {
	l, err := net.Listen("tcp", ":8888")
	if err != nil {
		fmt.Println("Listen error : ", err.Error())
		return
	}

	for {
		c, err := l.Accept()
		if err != nil {
			fmt.Println("Accept error : ", err.Error())
			break
		}

		go handleConn(c)
	}
}

func handleConn(c net.Conn) {
	defer c.Close()
	for {
		ioutil.ReadAll(c)
	}
}
