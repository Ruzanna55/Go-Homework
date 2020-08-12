package main

import "fmt"

func main() {
	fmt.Println(automat("0010001"))
}

const q0 = 0 // no input
const q1 = 1 // the number of '1' is 1
const q2 = 2 // the number of '1' is 0
const q3 = 3 // wrong input

func automat(input string) bool {
	var q = q0

	for i := 0; i < len(input); i++ {
		x := input[i]
		switch q {
		case q0:
			if x == '1' {
				q = q1

			} else {
				q = q2
			}

		case q1:
			if x == '1' {
				q = q3
			} else {
				q = q1
			}
		case q2:
			if x == '1' {
				q = q1

			} else {
				q = q2
			}
			// q = q3 ;nowhere to go

		}

	}
	return q == q1
}
