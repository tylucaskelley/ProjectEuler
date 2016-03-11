// finds sum of all multiples of 3 and 5 under 1000

#include <stdio.h>

int main(int argc, char** argv) {
    int sum = 0;
    int i;

    for (i = 1; i < 1000; i++) {
        if (i % 3 == 0 || i % 5 == 0) {
            sum += i;
        }
    }

    printf("%d\n", sum);

    return 0;
}
