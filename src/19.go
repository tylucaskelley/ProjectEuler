//Counting Sundays

package main

import (
    "fmt"
    "math"
)

func main() {
    fmt.Println(range_months(0,1901,2000))
}

func weekday(year, month, day int) int {
    d := day
    m := float64((month - 3) % 12 + 1)
    var y int
    
    if m > 10 {
        y = year - 1
    } else {
        y = year
    }
    
    y2 := y % 100
    c := (y - (y % 100)) / 100
    x := math.Floor((2.6 * m) - 0.2)
    
    w := ((d + int(x) + y2 + y2/4 + c/4 - 2*c)) % 7
    return w;
}

func range_months(day, start, end int) int {
    total := 0
    for year := start; year < end; year++ {
        for month := 1; month <= 12; month++ {
            if weekday(year, month, 1) == day {
                total++;   
            }
        }
    }
    
    return total;
}