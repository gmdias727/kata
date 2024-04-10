package ex1480.java;

import java.util.Arrays;

class Main {
    public static void main(String[] args) {
        int[] numbers = new int[]{1, 3, 5, 7, 9};
        runningSum(numbers);
    }

    public static void runningSum(int[] nums) {
        for (int i = 1; i < nums.length; i++) {
            nums[i] += nums[i - 1];
        }
        System.out.println(Arrays.toString(nums));
    }
}