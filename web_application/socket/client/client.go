package main

import (
	"fmt"
	"net"
	"os"
)

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error : %s", err.Error())
		os.Exit(1)
	}
}

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "usage :%s hoost:prot ", os.Args[0])
		os.Exit(1)
	}

	service := os.Args[1]
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	fmt.Println("tcpAddr : ", tcpAddr.String())
	checkError(err)
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	fmt.Println("connect succeed")
	checkError(err)
	_, err = conn.Write([]byte("HEAD / HTTP/1.0\r\n\r\n"))
	fmt.Println("write succeed")
	checkError(err)
	defer conn.Close()
	// result, err := ioutil.ReadAll(conn)
	// fmt.Println("read succeed")
	// checkError(err)
	// fmt.Println(string(result))
	os.Exit(0)
}
