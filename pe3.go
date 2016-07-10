package main

import "fmt"

func largestprime(n int) int {

	largest := 0

	for i := 3; i < n; i += 2 {
		if 600851475143 % i == 0 {
			largest = i
		}
	}
	return largest
}

func main() {

	z := largestprime(30000000)
	fmt.Println(z)
}
	
