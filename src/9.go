//Pythagorean triplet such that a + b + c = 1000

package main

import (
    "fmt"
    "math"
)

func main() {
    fmt.Println(pyProduct(1000));
}

func pyProduct(max int) int{
    for a := 0; a <= max/3; a++ {
        for b := a; b < max/2; b++ {
            c := max - a - b
            if isTriplet(a,b,c) {
                return a * b * c
            }
        }
    }
    return -1
}

func isTriplet(a,b,c int) bool {
    triplet := false
    if math.Pow(float64(a),2) + math.Pow(float64(b),2) == math.Pow(float64(c),2) {
        triplet = true
    } 
    return triplet
}