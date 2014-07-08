// Problem Four: Largest Palindrome from product of 3-digit integers

#include <iostream>
using namespace std;

//Define constants
static int MAX = 998001;
static int MIN = 10000;

//Define functions
int reverse_int(int num);
int largest_palindrome();

int main(int argc, char *argv[]) {
  int largest = largest_palindrome();
  cout << largest << endl;
  return 0;
} 

int reverse_int(int num) {
  int new_num = 0;
  while (num > 0) {
    new_num = new_num * 10 + (num % 10);
    num = num / 10;
  }
  return new_num;
}

int largest_palindrome() {
  int f1, f2;
  int large = 1;
  for (f1 = 999; f1 >= 100; f1--) {
    for (f2 = 999; f2 >= 100; f2--) {
      if ((f1 * f2 > large) && f1 * f2 == reverse_int(f1 * f2)) {
	large = f1 * f2;
      }
    }
  }
  return large;
}
