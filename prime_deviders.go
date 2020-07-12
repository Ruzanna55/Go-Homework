package main

import "fmt"

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
	fmt.Println("The program  outputs the prime dividers of the number.")
	fmt.Println("Enter the number.")
	n, err := fmt.Scan(&num)
	if n < 1 {
		fmt.Println("Error entering data", err)
		return
	}
	if !(num > 0) {
		fmt.Println("Input data error, required natural number")
		return
	}
	fmt.Print(num, "->")
	dividers(num)
	return
}
func dividers(num int) {
	a := 0
	k := 0
	if num == 1 {
		fmt.Print("1 doesn't have prime deviders")
		return
	}
	for i := 2; num > 1; i++ {
		if prime(i) {
			for num%i == 0 {
				a = i
				k++
				num = num / i
			}
			if k != 0 {
				fmt.Print(a, "^", k, "*")
			}
		}

		k = 0
	}
	return

}
