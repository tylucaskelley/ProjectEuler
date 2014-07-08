//Sum of Digits in 100!
package main
import (
    "fmt"
    "math/big"
    "strings"
    "strconv"
)

func main() {    
    x := big.NewInt(100)
    s := fact(x).String()
    
    sum := 0
    arr := strings.Split(s, "")
    
    for i := 0; i < len(arr); i++ {
        x,_ := strconv.Atoi(arr[i])
        sum += x
    }
    
    fmt.Println(sum)
}

func fact(n *big.Int) (r *big.Int) {
    r = new(big.Int)
    switch n.Cmp(&big.Int{}) {
        case -1,0:
            r.SetInt64(1)
        default:
            r.Set(n)
            var one big.Int
            one.SetInt64(1)
            r.Mul(r, fact(n.Sub(n, &one)))
    }
    
    return
}