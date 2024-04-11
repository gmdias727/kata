package ex412.java;

import java.util.ArrayList;
import java.util.List;

public class Main {
    public static void main(String[] args) {
        System.out.println(fizzFizz(5));
    }

    public static List<String> fizzFizz(int n) {
        List<String> answer = new ArrayList<>(n);

        for (int i = 1; i <= n; i++) {
            boolean divisibleBy3 = i % 3 == 0;
            boolean divisibleBy5 = i % 5 == 0;

            if (divisibleBy3 && divisibleBy5) {
                answer.add("FizzBuzz");
            } else if (divisibleBy3) {
                answer.add("Buzz");
            } else if (divisibleBy5) {
                answer.add("Fizz");
            } else {
                answer.add(String.valueOf(i));
            }
        }
        return answer;
    }
}