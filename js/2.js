// prints sum of all even-valued fibonacci numbers up to 4000000

var a = 1;
var b = 2;
var sum = 0;

while (a <= 4000000) {
    if (a % 2 == 0) {
        sum += a;
    }

    var temp = a;
    a = b;
    b = temp + b;
}

console.log(sum);
