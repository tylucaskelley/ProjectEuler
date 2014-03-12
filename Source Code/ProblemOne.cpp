//Problem #1: Sum of multiples of 3 or 5 under 1000
#include <iostream>
using namespace std;
const unsigned int MAX = 1000;

int main() {
  unsigned int sum = 0;
  for (unsigned int i = 0; i < MAX; ++i) {
    if (i % 3 == 0 || i % 5 == 0) {
      sum += i;
    }
  }
  cout << "Sum: " << sum << endl;
  return 0;
}
