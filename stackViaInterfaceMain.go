package main

import (
	"fmt"

	"./stackviainterface"
)

type stack interface {
	Push(int) bool
	Pop() (int, bool)
}

func main() {
	var s stack
	s = new(stackviainterface.ArrayStack)
	s.Push(10)
	s.Push(110)
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())

	return
}
