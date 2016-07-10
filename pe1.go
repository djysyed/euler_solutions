package main

import "fmt"

func three(n int) int {
	total := 0

	for i := 1; i < n; i++ {
		if i%3 == 0 {
			total += i
		}
	}
	return total
}

func five(n int) int {
	total := 0

	for i := 1; i < n; i++ {
		if i%5 == 0 {
			total += i
		}
	}
	return total
}

func fifteen(n int) int {
	total := 0

	for i := 1; i < n; i++ {
		if i%15 == 0 {
			total += i
		}
	}
	return total
}
func main() {
	z := three(20) + five(20) - fifteen(20)
	fmt.Println(z)
}
