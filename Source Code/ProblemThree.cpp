//Problem #3: Largest Prime Factor

#include <iostream>

using namespace std;

int main() {
	long long int num = 600851475143;

	for (long i = 2; i < num; i++) {
		if (num % i == 0) {
			num /= i;
			i -= 1;
		}
	}

	cout << "Largest Prime Factor: " << num << endl;
	return 0;
}