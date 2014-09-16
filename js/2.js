function fibSum(n) {
    var a = 1;
    var b = 2;
    var sum = 0;
    
    while (a <= n) {
        if (a % 2 == 0) {
            sum += a;
        }

        var temp = a;
        a = b;
        b = temp + b;
    }

    return sum;
};

console.log(fibSum(4000000));
