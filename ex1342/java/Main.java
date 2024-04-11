package ex1342.java;

public class Main {
    public static void main(String[] args) {
        System.out.println(numberOfSteps(20));
    }
    public static int numberOfSteps(int num) {
        int steps = 0;
        for (int i = 0; i < num; i++) {
            if (num % 2 == 0) {
                num /= 2;
            } else {
                num -= 1;
            }
            steps++;
        }
        return steps;
    }
}