// go run main.go myFile.txt
// go run main.go main.go won't work
// instead, do the following
// go build main.go
// ./main main.go

package main

import (
	"fmt"
	"io"
	"os"
)

type myPrinter struct{}

func main() {
	fmt.Println(os.Args[1])
	f, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	mp := myPrinter{}
	io.Copy(mp, f)

	// bs, err := os.ReadFile("myFile.txt")
	// if err != nil {
	// 	panic(err)
	// }
	// println(string(bs))
}

func (myPrinter) Write(bs []byte) (int, error) {
	fmt.Println("This is my custom writer!")
	fmt.Println(string(bs))
	fmt.Println("Just wrote this many bytes:", len(bs))
	return len(bs), nil
}
