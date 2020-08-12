package main

import "fmt"

func main() {
	fmt.Println(automat("00100"))
}

const q0 = 0 // no input
const q1 = 1 // the input number is '1'
const q2 = 2 // the input number is '0'

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
				q = q1
			} else {
				q = q2
			}
		case q2:
			if x == '1' {
				q = q1

			} else {
				q = q2
			}

		}

	}
	return q == q1
}
