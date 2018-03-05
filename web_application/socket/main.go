package main

import "fmt"
import "net"
import "os"
import "time"

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error : %s  \n", err.Error())
		os.Exit(1)
	}
}

func main() {
	// if len(os.Args) != 2 {
	//     fmt.Fprintf(os.Stderr, "Usage : %s ip addr \n", os.Args[0])
	//     os.Exit(1)
	// }
	//
	// name := os.Args[1]
	// addr := net.ParseIP(name)
	// if addr == nil {
	//     fmt.Println("invalid address")
	// } else {
	//     fmt.Println("the address is ", addr.String())
	// }
	// os.Exit(0)

	service := ":12000"
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	checkError(err)
	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)
	for {
		fmt.Println("listener Accept before")
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("listener err : ", err.Error())
			continue
		}

		fmt.Println("conn write")
		daytime := time.Now().String()
		conn.Write([]byte(daytime))
		conn.Close()
	}

}
