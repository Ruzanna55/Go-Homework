package main

import (
	"fmt"

	"./stackarray"
)

func main() {
	fmt.Println(stackarray.Push(3))
	fmt.Println(stackarray.Push(10))
	fmt.Println(stackarray.Push(5))
	fmt.Println(stackarray.Pop())

	return

}
