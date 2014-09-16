public class Two {
    public static void main(String[] args) {
        System.out.printf("%d\n", fibSum(4000000));
    }

    public static int fibSum(int max) {
        int temp;
        int sum = 0;
        int a = 1;
        int b = 2;

        while (a <= max) {
            if (a % 2 == 0) {
                sum += a;
            }

            temp = a;
            a = b;
            b = temp + a;
        }

        return sum;
    }
}
