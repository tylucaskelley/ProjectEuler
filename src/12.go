//Highly divisible triangle numbers
package main

import (
    "math"
    "fmt"
)

func main() {
    highNum, highFactors, i := 0, 0, 1
        for i <= int(^uint(0) >> 1) {
            if highFactors > 500 {
                break
            }
            
            sum, j := 0, 1
            for j <= i {
                sum += j
                j++
            }
            
            current := divisibility(sum)
            if current > highFactors {
                highFactors = current
                highNum = sum
            }            
            i++
        } 
    
    fmt.Println(highNum)
}

func divisibility(n int) int {
    x := int(math.Sqrt(float64(n)))
    count, i := 0, 1
    for i < x {
        if n%i == 0 {
            count+=2
        }
        i++
    }
    return count
}