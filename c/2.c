// prints sum of all even-valued fibonacci numbers up to 4000000

#include <stdio.h>

int fibSum(int);

int main(int argc, char** argv) {
    int a = 1;
    int b = 2;
    int sum = 0;
    int temp;

    while (a < 4000000) {
        if (a % 2 == 0) {
            sum += a;
        }

        temp = a;
        a = b;
        b = temp + b;
    }

    printf("%d\n", sum);
    return 0;
}
