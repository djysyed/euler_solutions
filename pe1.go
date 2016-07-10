package main

import "fmt"

func sum(n int) int {
	total := 0

	for i := 1; i < n; i++ {
		if i%3 == 0 || i%5 == 0 {
			total += i
		}
	}
	return total
}

func main() {
	z := sum(1000)
	fmt.Println(z)
}
