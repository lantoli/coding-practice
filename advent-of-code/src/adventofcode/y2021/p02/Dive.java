package adventofcode.y2021.p02;

import java.io.InputStream;
import java.util.Scanner;

public class Dive {

    public static void main(String[] args) throws RuntimeException {
        new Dive().run2();
    }

    // h: 2165, d: 933, mul: 2019945
    private void run1() {
        Scanner s = getScanner();
        int h = 0, d = 0;
        while (s.hasNext()) {
            String str = s.next();
            int num = s.nextInt();
            switch (str) {
                case "forward":
                    h += num;
                    break;
                case "down":
                    d += num;
                    break;
                case "up":
                    d -= num;
                    break;
            }
        }
        System.out.printf("h: %d, d: %d, mul: %d\n", h, d, h * d);
    }

    // h: 2165, d: 738712, mul: 1599311480
    private void run2() {
        Scanner s = getScanner();
        int h = 0, d = 0, a = 0;
        while (s.hasNext()) {
            String str = s.next();
            int num = s.nextInt();
            switch (str) {
                case "forward":
                    h += num;
                    d += num * a;
                    break;
                case "down":
                    a += num;
                    break;
                case "up":
                    a -= num;
                    break;
            }
        }
        System.out.printf("h: %d, d: %d, mul: %d\n", h, d, h * d);
    }


    private Scanner getScanner() {
        InputStream in = getClass().getResourceAsStream("input.txt");
        return new Scanner(in);
    }
}
