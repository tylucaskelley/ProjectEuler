//Sum - Square Difference
package main
import (
    "fmt"
    "math"
)

const (
    MIN float64 = 1
    MAX float64 = 100
)

func main() {
    i := MIN
    sumSquares := 0.0
    sumInts := 0.0
    
    for i <= MAX {
        sumSquares += math.Pow(i,2.0) 
        sumInts += i
        i += 1
    }
    
    squareSums := math.Pow(sumInts,2.0)
    diff := squareSums - sumSquares
    
    fmt.Printf("%f", diff)
}