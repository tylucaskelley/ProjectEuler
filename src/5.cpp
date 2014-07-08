//Problem Five: Smallest number that can be divided by all integers 1 - 20

#include <iostream>
using namespace std;

static int START = 2520;

bool smallest_multiple(int x);

int main(int argc, char *argv[]) {
  int i = START;
  while (true) {
    if (smallest_multiple(i)) {
      break;
    }
    i += 20;
  }
  cout << i << endl;
}

bool smallest_multiple(int x) {
  int i = 11;
  for (i; i <= 20; i++) {
    if (x % i != 0) {
      return false;
    }
  }
  return true;
}
