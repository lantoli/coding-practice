package adventofcode.y2021.p17;

import java.io.InputStream;
import java.time.Duration;
import java.time.Instant;
import java.util.ArrayList;
import java.util.Scanner;

public class Shot {

    public static void main(String[] args) throws RuntimeException {
        new Shot().process();
    }

    // 9180, 3767
    private void process() {
        Instant one = Instant.now();
        Scanner s = getScanner();

        // var res = run1(150, 193, -136, -86);
        // var res = run2(20, 30, -10, -5);
        var res = run2(150, 193, -136, -86);
        System.out.println(res);

        Instant two = Instant.now();
        Duration dur = Duration.between(one, two);
        System.out.println(dur);
    }

    private int run1(int x1, int x2, int y1, int y2) {
        var ret = 0;
        for (int xx = 0; xx <= XMAX; xx++)
            next: for (int yy = -YMAX; yy <= YMAX; yy++) {
                int maxy= 0, x = 0, y = 0, xinc = xx, yinc = yy;
                while (x <= x2 && y >= y2) {
                    x += xinc;
                    y += yinc;
                    maxy = Math.max(maxy, y);
                    if (xinc > 0) xinc--;
                    yinc--;
                    if (x >= x1 && x <= x2 && y >= y1 && y <= y2) {
                        ret = Math.max(ret, maxy);
                        continue next;
                    }
                }
            }
        return ret;
    }

    private int run2(int x1, int x2, int y1, int y2) {
        var ret = 0;
        var list = new ArrayList<int[]>();
        for (int xx = 0; xx <= XMAX; xx++)
            next: for (int yy = -YMAX; yy <= YMAX; yy++) {
                int x = 0, y = 0, xinc = xx, yinc = yy;
                while (x <= x2 && y >= y1) {
                    x += xinc;
                    y += yinc;
                    if (xinc > 0) xinc--;
                    yinc--;
                    if (x >= x1 && x <= x2 && y >= y1 && y <= y2) {
                        ret++;
                        list.add(new int[] {xx, yy});
                        continue next;
                    }
                }
            }
        return ret;
    }

    private final static int XMAX = 10000;
    private final static int YMAX = 10000;
    
    private Scanner getScanner() {
        InputStream in = getClass().getResourceAsStream("input.txt");
        return new Scanner(in);
    }
}
