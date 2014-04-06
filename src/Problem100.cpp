/*
 * Arranged Probability
 */

#include <iostream>
#include <cmath>
using namespace std;

class Problem100 {
		long n;
	public:
		long num_blues();
		void set_n(long);
};

void Problem100::set_n(long num) {
	n = num;
}

long Problem100::num_blues() {
	long result;

	long init_b = 15;
	long init_n = 21;

	while (init_n < n) {
		long temp_b = (3*init_b) + (2*init_n) - 2;
		long temp_n = (4*init_b) + (3*init_n) - 3;

		init_b = temp_b;
		init_n = temp_n;
	}
	result = init_b;

	return result;
}

int main() {
	Problem100 p;
	p.set_n(1000000000000);
	cout << "Number of Blue Disks: " << p.num_blues() << endl;
	return 0;
}