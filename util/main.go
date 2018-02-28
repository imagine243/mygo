package main

import (
	"fmt"
	"os"
)

func main() {
	// fmt.Println("hello util")
	// fmt.Print(rand.Intn(100), ",")
	// fmt.Print(rand.Intn(100))
	//
	// fmt.Println()
	//
	// fmt.Println(rand.Float64())
	//
	// fmt.Print(rand.Float64()*5+5, ",")
	// fmt.Print(rand.Float64()*5 + 5)
	// fmt.Println()
	//
	// s1 := rand.NewSource(time.Now().UnixNano())
	// r1 := rand.New(s1)
	//
	// fmt.Print(r1.Intn(100), ",")
	// fmt.Println(r1.Intn(100))
	//
	// s2 := rand.NewSource(42)
	// r2 := rand.New(s2)
	// fmt.Print(r2.Intn(100), ",")
	// fmt.Println(r2.Intn(100))
	//
	// s3 := rand.NewSource(42)
	// r3 := rand.New(s3)
	// fmt.Print(r3.Intn(100), ",")
	// fmt.Println(r3.Intn(100))

	// dat, err := ioutil.ReadFile("/tmp/dat")
	// check(err)
	// fmt.Print(string(dat))

	// dateCmd := exec.Command("date")
	// dateOut, err := dateCmd.Output()
	// check(err)
	// fmt.Println(">date")
	// fmt.Println(string(dateOut))
	//
	// grepCmd := exec.Command("grep", "hello")
	// grepIn, _ := grepCmd.StdinPipe()
	// grepOut, _ := grepCmd.StdoutPipe()
	// grepCmd.Start()
	// grepIn.Write([]byte("hello grep \n goodbye grep"))
	// grepIn.Close()
	// grepBytes, _ := ioutil.ReadAll(grepOut)
	// grepCmd.Wait()
	// fmt.Println(string(grepBytes))
	//
	// lsCmd := exec.Command("bash", "-c", "ls -a -l -h")
	// lsOut, err := lsCmd.Output()
	// check(err)
	// fmt.Println(">ls -a -l -h")
	// fmt.Println(string(lsOut))

	// binary, lookErr := exec.LookPath("ls")
	// check(lookErr)
	//
	// args := []string{"ls", "-a", "-l", "-h"}
	// env := os.Environ()
	//
	// execErr := syscall.Exec(binary, args, env)
	// check(execErr)

	// sigs := make(chan os.Signal, 1)
	// done := make(chan bool, 1)
	//
	// signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	//
	// go func() {
	//     sig := <-sigs
	//     fmt.Println()
	//     fmt.Println(sig)
	//     done <- true
	// }()
	//
	// fmt.Println("awaiting signal")
	// <-done
	// fmt.Println("exiting")

	defer fmt.Println("!")
	os.Exit(3)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
