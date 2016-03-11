# finds sum of all multiples of 3 and 5 under 1000

sum = 0

(1..999).each do |x|
    sum += x if (x % 3 == 0) || (x % 5 == 0)
end

puts sum
