package main

import (
	"fmt"
)

func prime(num int) bool {
	a := true
	for i := 2; i*i <= num; i++ {
		if num%i == 0 {
			a = false
			break
		}
	}
	if num == 1 {
		a = false
	}

	return a
}
func main() {
	var num int
	var answer bool
	fmt.Println("The program checks that the entered number is prime.")
	fmt.Println("Enter the number.")
	n, err := fmt.Scan(&num)
	if n < 1 {
		fmt.Println("Error entering data", err)
		return
	}
	if !(num > 0) {
		fmt.Println("Input data error, required prime number")
		return
	}
	answer = prime(num)
	fmt.Println(answer)
	if answer == true {
		fmt.Println("The number is prime")
	} else {
		fmt.Println("The number isn't prime")
	}
	return
}
