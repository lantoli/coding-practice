package adventofcode.y2021.p03;

import java.io.InputStream;
import java.util.ArrayList;
import java.util.Scanner;

public class Binary {

    public static void main(String[] args) throws RuntimeException {
        new Binary().run2();
    }

    // gamma: 2484, epsilon: 1611, mul: 4001724
    private void run1() {
        Scanner s = getScanner();
        int total = 0;
        int[] bits = new int[BITS];
        while (s.hasNext()) {
            total++;
            String str = s.nextLine();
            for (int i = 0; i < BITS; i++) {
                if (str.charAt(i) == '1') bits[i]++;
            }
        }
        int gamma = 0;
        for (int i = 0; i < BITS; i++) {
            if (bits[BITS - i - 1] > total/2) gamma += 1 << i;
        }
        int epsilon = (1 << BITS) - gamma - 1;

        System.out.printf("gamma: %d, epsilon: %d, mul: %d\n", gamma, epsilon, gamma * epsilon);
    }

    // oxygen: 2545, co2: 231, mul: 587895
    private void run2() {
        Scanner s = getScanner();
        var list = new ArrayList<int[]>();
        while (s.hasNext()) {
            String str = s.nextLine();
            int[] arr = new int[BITS+1];
            for (int i = 0; i < BITS; i++) if (str.charAt(i) == '1') {
                arr[i] = 1;
                arr[BITS] += 1 << (BITS - i -1);
            }
            list.add(arr);
        }
        int oxygen = getNumber(list, 1, 0 );
        int co2 = getNumber(list, 0, 1 );
        System.out.printf("oxygen: %d, co2: %d, mul: %d\n", oxygen, co2, oxygen * co2);
    }

    private int getNumber(ArrayList<int[]> listParam, int first, int second) {
        var list = new ArrayList<>(listParam);
        int ret = -1;
        for (int i=0; i<BITS; i++) {
            int ii = i;
            int ones = list.stream().filter(arr -> arr[ii] == 1).toArray().length;
            int number = (2 * ones >= list.size()) ? first : second;
            list.removeIf(arr -> arr[ii] != number);
            if (list.size() == 1) {
                ret = list.get(0)[BITS];
                break;
            }
        }
        return ret;
    }

    final private static int BITS = 12;

    private Scanner getScanner() {
        InputStream in = getClass().getResourceAsStream("input.txt");
        return new Scanner(in);
    }
}
