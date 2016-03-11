#!/usr/bin/env bash
# prints sum of all even-valued fibonacci numbers up to 4000000

sum=0
a=1
b=2

while (($a < 4000000)); do
    if (($a % 2 == 0)); then
        sum=$((sum + a))
    fi

    temp=$a
    a=$b
    b=$((temp + b))
done

echo $sum
