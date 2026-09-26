package adventofcode.y2021.p06;

import java.io.InputStream;
import java.util.Arrays;
import java.util.HashMap;
import java.util.List;
import java.util.Scanner;
import static java.util.stream.Collectors.toList;

public class Lanternfish {

    public static void main(String[] args) throws RuntimeException {
        new Lanternfish().process();
    }

    private void process() {
        Scanner s = getScanner();
        var list = Arrays.stream(s.nextLine().split(",")).map(Integer::valueOf).collect(toList());
        run2(list);
    }

    // 363101
    private void run1(List<Integer> list) {
        for (int day = 0; day < 80; day++) {
            for (int i = list.size() - 1; i >= 0; i--) {
                int cur = list.get(i);
                if (cur == 0) {
                    list.set(i, 6);
                    list.add(8);
                } else {
                    list.set(i, cur - 1);
                }
            }
        }
        System.out.printf("total: %d\n", list.size());
    }

    // 26984457539
    private void run2(List<Integer> list) {
        long[] ages = new long[9];
        for (int elm : list) ages[elm]++;

        for (int day = 0; day < 256; day++) {
            long zero = ages[0];
            for (int i=0; i<8; i++) {
                ages[i] = ages[i+1];
            }
            ages[8] = zero;
            ages[6] += zero;
        }
        System.out.printf("total: %d\n", Arrays.stream(ages).sum());
    }

    private Scanner getScanner() {
        InputStream in = getClass().getResourceAsStream("input.txt");
        return new Scanner(in);
    }
}
