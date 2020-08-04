
package stackviainterface

import "fmt"

type ArrayStack struct {
	arr [4]int
	top int
}

var n = 4

// Push the element to the stack
func (s *ArrayStack) Push(val int) bool {
	if s.top >= n-1 {
		fmt.Println("Stack Overflow")
		return false
	}
	s.top++
	s.arr[s.top] = val
	fmt.Println("The pushed element is ", s.arr[s.top])
	return true

}

// Pop the element from the stack
func (s *ArrayStack) Pop() (int, bool) {

	if (s.top) <= 0 {
		fmt.Println("Stack Underflow")
		return -1, false
	}
	fmt.Println("The popped element is ")
	a := s.arr[s.top]
	(s.top)--
	return a, true

}

