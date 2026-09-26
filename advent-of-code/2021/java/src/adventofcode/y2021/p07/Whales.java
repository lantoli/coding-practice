package adventofcode.y2021.p07;

import java.io.InputStream;
import java.util.Arrays;
import java.util.Scanner;

public class Whales {

    public static void main(String[] args) throws RuntimeException {
        new Whales().process();
    }

    private void process() {
        Scanner s = getScanner();
        var list = Arrays.stream(s.nextLine().split(",")).mapToInt(Integer::parseInt).toArray();
        int min = Arrays.stream(list).min().getAsInt();
        int max = Arrays.stream(list).max().getAsInt();
        run(list, min, max);
    }

    private void run(int[] list, int min, int max) {
        int ret = Integer.MAX_VALUE;

        for (int i = min; i<=max; i++) {
            int fuel = 0;
            for (int n=0; n<list.length; n++) fuel += dist2(list[n], i);
            ret = Math.min(ret, fuel);
        }

        System.out.printf("total: %d\n", ret);
    }

    // 356922
    private int dist1(int a, int b) {
        return Math.abs(a - b);
    }

    // ...
    private int dist2(int a, int b) {
        int total = Math.abs(a - b);
        return total * (total + 1) / 2;
    }

    private Scanner getScanner() {
        InputStream in = getClass().getResourceAsStream("input.txt");
        return new Scanner(in);
    }
}
