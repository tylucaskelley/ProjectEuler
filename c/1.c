#include <stdio.h>
#include <stdbool.h>

bool isMult(int);

int main(int argc, char** argv) {
    int sum = 0;
    int i;

    for (i = 1; i < 1000; i++) {
        if (isMult(i)) {
            sum += i;
        }
    }

    printf("%d\n", sum);
    return 0;
}

bool isMult(int x) {
    if (x % 3 == 0 || x % 5 == 0) {
        return true;
    }
    return false;
}
