#include <stdio.h>

int fibSum(int);

int main(int argc, char** argv) {
    printf("%d\n", fibSum(4000000));
    
    return 0;
}

int fibSum(int max) {
    int a = 1;
    int b = 2;
    int sum = 0;
    int temp;

    while (a < max) {
        if (a % 2 == 0) {
            sum += a;   
        }
        temp = a;
        a = b;
        b = temp + b;
    }

    return sum;
}
