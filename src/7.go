package main

import "fmt"

func main() {
	x := 2
	count := 1
	for count < 10001 {
		x++
		if isPrime(x) {
			count++
		}
	}

	fmt.Println(x)
}

func isPrime(x int) bool {
	i := 2
	for i < x {
		if x%i == 0 {
			return false
		}
		i++
	}
	return true
}
