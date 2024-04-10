package ex1672.java;

public class Main {
    public static void main(String[] args) {
        int[][] array = { {1,2,3}, {3,2,1} };
        maximumWealth(array);
    }
    public static int maximumWealth(int[][] accounts) {
        for (int i = 0; i < accounts.length; i++) {
            int sum = 0;
            for (int j = 0; j < accounts[i].length; j++) {
                sum += accounts[i][j];
            }
            return sum;
        }
    }
}
