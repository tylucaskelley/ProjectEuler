//Max sum in a triangle of numbers
package main

import (
    "fmt"
    "bufio"
    "io"
    "log"
    "os"
    "strings"
    "math"
    "strconv"
)

func main() {
    //read from the file
    f, err := os.Open("../files/p18.txt")
    if err != nil {
        log.Fatal(err)
    }
    
    s := make([]string, 15)
    reader := bufio.NewReader(f)
    for i:= 0; ; i++ {
        line, isPrefix, err := reader.ReadLine()
        if err == io.EOF {
            break
        } else if err != nil {
            log.Fatal(err)
        } else if isPrefix {
            log.Fatal("Line too long", f.Name())
        } else {
            //everything is good
            s[i] = string(line)
        }
    }
    
    //it seemed like a good way to make a 2d array at the time...
    var triangle [15][15]string
    for i := 0; i < len(s); i++ {
        arr := [15]string{
            "0",
            "0",
            "0",
            "0",
            "0",
            "0",
            "0",
            "0",
            "0",
            "0",
            "0",
            "0",
            "0",
            "0",
            "0",
        }
        temp := strings.Split(s[i], " ")
        for j := 0; j < len(temp); j++ {
            arr[j] = temp[j]   
        }
        
        triangle[i] = arr
    }
    
    var nums [15][15]int
    for i := 0; i < len(triangle); i++ {
        for j := 0; j < len(triangle); j++ {
            a, err := strconv.Atoi(triangle[i][j])
            if err != nil {
                log.Fatal(err) 
            }
            
            nums[i][j] = a
        }
    }
    
    //the actual solution (screw Go for not having 2d arrays)
    for i := 13; i >= 0; i-- {
        for j := 0; j <= i; j++ {
            nums[i][j] += int(math.Max(float64(nums[i+1][j]), float64(nums[i+1][j+1])))
        }
    }
    
    fmt.Println(nums[0][0])
}   