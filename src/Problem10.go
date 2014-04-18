package main

import (
	"fmt"
)

func main() {
	const MAX int = 2000000
	fmt.Println(sumPrimes(MAX))
}

func sumPrimes(max int) int {
	total := 1

	i := 0
	for i < max {
		if isPrime(i) {
			total += i
            fmt.Println(total)
		}
		i++
	}

	return total
}

func isPrime(x int) bool {
	if x%2 == 0 || x%3 == 0 || x%5 == 0 || x%7 == 0 || x%11 == 0 {
		return false
	}

	i := 3
	for i < x {
		if x%i == 0 {
			return false
		}
		i += 2
	}

	return true
}
