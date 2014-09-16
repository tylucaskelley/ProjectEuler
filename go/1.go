package main

import "fmt"

func isMult(i int) bool {
    return i % 3 == 0 || i % 5 == 0
}

func main() {
    sum := 0
    for i := 0; i < 1000; i++ {
        if isMult(i) {
            sum += i
        }
    }
    fmt.Printf("%d\n", sum)
}
