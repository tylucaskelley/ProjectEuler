# Project Euler Solutions

**Ty-Lucas Kelley**

---

### Goal

To solve the same math problems in 5 different languages, simply to see how the solutions compare in terms of speed, readability, and other factors.

### Languages used

* C
* Go
* JavaScript (NodeJS)
* Python
* Java

### Current solutions

* 1: Multiples of 3 and 5
* 2: Even Fibonacci numbers

### Problem 1 in all 5 languages

C

```c
    #include <stdio.h>
    #include <stdbool.h>

    bool isMult(int);

    int main(int argc, char** argv) {
        int sum = 0;
        int i;

        for (i = 1; i < 1000; i++) {
            if (isMult(i)) {
                sum += i;
            }
        }

        printf("%d\n", sum);
        return 0;
    }

    bool isMult(int x) {
        if (x % 3 == 0 || x % 5 == 0) {
            return true;
        }
        return false;
    }
```

Python

```python
    print sum([i for i in xrange(1000) if i % 3 == 0 or i % 5 == 0])
```

Go

```go
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
```

JavaScript

```js
    function multiples(n) {
        var sum = 0;

        for (var i = 0; i < n; i++) {
            if (i % 3 === 0 || i % 5 === 0) {
                sum += i;
            }
        }

        return sum;
    }

    console.log(multiples(1000));
```

Java

```java
    public class One {
        public static void main(String[] args) {
            System.out.printf("%d\n", multiples(1000));        
        }

        public static int multiples(int n) {
            int sum = 0;

            for (int i = 0; i < n; i++) {
                if (i % 3 == 0 || i % 5 == 0) {
                    sum += i;
                }
            }

            return sum;
        }
    }
```
