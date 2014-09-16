#!/usr/bin/python

def fib_sum(n):
    s, a, b = (0, 1, 2)
    while a <= n:
        if a % 2 == 0:
            s += a
        a, b = b, a + b
    return s

print fib_sum(4000000)
