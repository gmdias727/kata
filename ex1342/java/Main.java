package ex1342.java;

public class Main {
    public static void main(String[] args) {
        int com8 = 8;
        int com14 = 14;
        int com123 = 123;
        System.out.printf("com 8 steps deve ser 4, resultado: %d \n", numberOfSteps(com8));
        System.out.printf("com 14 steps deve ser 6, resultado: %d \n", numberOfSteps(com14));
        System.out.printf("com 123 steps deve ser 12, resultado: %d \n", numberOfSteps(com123));
    }
    public static int numberOfSteps(int num) {
        int steps = 0;
        while (num > 0) {
            if (num % 2 == 0) {
                num /= 2;
            } else {
                num--;
            }
            steps++;
        }
        return steps;
    }
}
