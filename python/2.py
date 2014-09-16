#!/usr/bin/python

def fib_sum(n):
    sum, a, b = (0, 1, 2)
    while a <= n:
        if a % 2 == 0:
            sum += a
        a, b = b, a + b
    return sum

print fib_sum(4000000)
