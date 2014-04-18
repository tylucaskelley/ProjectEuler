//Longest Collatz Sequence for n under 1000000
package main

import (
    "fmt"
    "os"
)

func main() {
    i, high := 2, 0
    for i < 1000000 {
        if collatz(i) > high {
            high = collatz(i)
            fmt.Println(i)
        }
        i++
    } 
}

func collatz(n int) int {
    if n < 2 {
        fmt.Println("n must be >= 2")
        os.Exit(2)
    }
    
    count := 1
    for n != 1 {
        if n%2 == 0 {
            n = n/2
        } else { 
            n = 3*n + 1
        }
        count++
    } 
    return count
}