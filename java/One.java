// finds sum of all multiples of 3 and 5 under 1000

public class One {
    public static void main(String[] args) {
        System.out.printf("%d\n", multiples(1000));
    }

    public static int multiples(int n) {
        int sum = 0;

        for (int i = 0; i < n; i++) {
            if (i % 3 == 0 || i % 5 == 0) {
                sum += i;
            }
        }

        return sum;
    }
}
