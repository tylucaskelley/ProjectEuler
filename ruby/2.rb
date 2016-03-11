# prints sum of all even-valued fibonacci numbers up to 4000000

sum, a, b = 0, 1, 2

while a < 4000000 do
    sum += a if a % 2 == 0
    a, b = b, a + b
end

puts(sum)
