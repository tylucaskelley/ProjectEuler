//Power Digit Sum
package main

import (
    "fmt"
    "math"
    "strconv"
    "strings"
    //"os"
)

//Takes a 128-bit floating point number and represents it as a string
var s string = strconv.FormatFloat(math.Exp2(1000), 'f', 0, 64)

func main() {
    sum := 0
    arr := strings.Split(s, "")
    
    for i := 0; i < len(arr); i++ {
        x,_ := strconv.Atoi(arr[i])
        sum += x
    }
    
    fmt.Println(sum)
}