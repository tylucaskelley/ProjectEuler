//Problem #2: Even Fibonacci Numbers

#include <iostream>

using namespace std;

int main() {
	unsigned int a = 1;
	unsigned int b = 2;
	unsigned int sum = 0;

	while (a < 4000000) {
		b = a + b;
		a = b - a;
		if (a % 2 == 0) {
			sum += a;
		}
	}

	cout << "Sum of even Fibonacci numbers: " << sum << endl;
	return 0;
}