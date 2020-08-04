package main

import "fmt"

var stack [3]int
var n = 3
var top = -1

func main() {
	fmt.Println(push(9))

	fmt.Println(pop())
	fmt.Println(pop())

	fmt.Println(stackElement())
	/*fmt.Println(pop())*/
	return

}

// Push the element to the stack
func Push(val int) bool {
	if top >= n-1 {
		fmt.Println("Stack Overflow")
		return false
	}
	top++
	stack[top] = val
	fmt.Print(stack[top])
	return true
}

// Pop the element from the stack
func Pop() (int, bool) {
	if top <= -1 {
		fmt.Println("Stack Underflow")
		return -1, false
	}
	fmt.Println("The popped element is ")
	a := stack[top]
	top--
	return a, true

}
func stackElement() bool {
	if top >= 0 {
		fmt.Println("Stack elements are:")
		for i := top; i >= 0; i-- {
			fmt.Print(stack[i], " ")
		}
		return true
	}
	fmt.Println("Stack is empty")
	return false
}
