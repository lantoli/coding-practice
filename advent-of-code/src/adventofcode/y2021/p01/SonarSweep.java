package adventofcode.y2021.p01;

import java.io.InputStream;
import java.util.Scanner;

public class SonarSweep {

    public static void main(String[] args) throws RuntimeException {
        new SonarSweep().run2();
    }

    // 1616
    private void run1() {
        Scanner s = getScanner();
        int total = 0;
        int last = s.nextInt();
        while (s.hasNext()) {
            int next = s.nextInt();
            if (next > last) total++;
            last = next;
        }
        System.out.println(total);
    }

    // 1645
    private void run2() {
        Scanner s = getScanner();
        int total = 0;
        int a = s.nextInt(), b = s.nextInt(), c = s.nextInt();
        while (s.hasNext()) {
            int next = s.nextInt();
            if (next > a) total++;
            a = b; b = c; c = next;
        }
        System.out.println(total);
    }

    private Scanner getScanner() {
        InputStream in = getClass().getResourceAsStream("input.txt");
        return new Scanner(in);
    }
}
