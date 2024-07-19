package main

import "fmt"

type dog struct{}
type cat struct{}

func (d dog) speak() {
	fmt.Println("Woof Woof")
}

func (c cat) speak() {
	fmt.Println("Meow Meow")
}

type animal interface {
	speak()
}

func main() {
	animals := []animal{dog{}, cat{}}
	for _, animal := range animals {
		animal.speak()
	}
}
