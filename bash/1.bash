#!/usr/bin/env bash
# finds sum of all multiples of 3 and 5 under 1000

sum=0

for i in {1..999}; do
    if (($i % 3 == 0 || $i % 5 == 0)); then
        sum=$((sum + i))
    fi
done

echo $sum
