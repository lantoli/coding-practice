package adventofcode.y2021.p22;

import java.io.InputStream;
import java.time.Duration;
import java.time.Instant;
import java.util.ArrayList;
import java.util.BitSet;
import java.util.Collections;
import java.util.List;
import java.util.Scanner;
import java.util.TreeSet;
import static java.lang.Integer.parseInt;

public class Reboot {

    public static void main(String[] args) throws RuntimeException {
        new Reboot().process();
    }

    private void process() {
        Instant one = Instant.now();
        Scanner s = getScanner();
        var steps = new ArrayList<Step>();
        while (s.hasNext()) {
            var line = s.nextLine().split(" ");
            boolean on = line[0].equals("on");
            var coord = line[1].split(",");
            var xcoord = coord[0].substring(2).split("\\.\\.");
            var ycoord = coord[1].substring(2).split("\\.\\.");
            var zcoord = coord[2].substring(2).split("\\.\\.");
            steps.add(new Step(parse(xcoord[0]), parse(xcoord[1]), parse(ycoord[0]), parse(ycoord[1]),
                               parse(zcoord[0]), parse(zcoord[1]), on));
        }

        var ret = run2(steps);
        System.out.printf("RET: %d\n", ret);

        Instant two = Instant.now();
        Duration dur = Duration.between(one, two);
        System.out.println(dur);
    }

    private int parse(String coord) {
        return parseInt(coord) + NEG;
    }

    // 658691
    private int run1(List<Step> steps) {
        boolean[][][] grid = new boolean[SIZE][SIZE][SIZE];

        for (var step: steps) {
            for (int z = Math.max(0, step.z1); z <= Math.min(SIZE-1, step.z2); z++)
                for (int y = Math.max(0, step.y1); y <= Math.min(SIZE-1, step.y2); y++)
                    for (int x = Math.max(0, step.x1); x <= Math.min(SIZE-1, step.x2); x++)
                        grid[z][y][x] = step.on;
        }

        int total = 0;
        for (int z = 0; z < SIZE; z++)
            for (int y = 0; y < SIZE; y++)
                for (int x = 0; x < SIZE; x++)
                    if (grid[z][y][x]) total++;
        return total;
    }

    private long run2(List<Step> steps) {
        Collections.reverse(steps);
        var visited = new TreeSet<Long>();
        long total = 0;
        for (int i=0; i<steps.size(); i++) {
            System.out.printf("%d ", i);
            var step = steps.get(i);
        //for (var step: steps) {
            for (int z = step.z1; z <= step.z2; z++)
                for (int y = step.y1; y <= step.y2; y++)
                    for (int x = step.x1; x <= step.x2; x++) {
                        long pos = z * SIZE * SIZE + y * SIZE + x;
                        if (step.on && !visited.contains(pos)) {
                            total++;
                        }
                        visited.add(pos);
                    }
        }
        return total;
    }

    private void minMax(List<Step> steps) {
        int minx = Integer.MAX_VALUE, miny = Integer.MAX_VALUE, minz = Integer.MAX_VALUE;
        int maxx = Integer.MIN_VALUE, maxy = Integer.MIN_VALUE, maxz = Integer.MIN_VALUE;
        for (var step: steps) {
            minx = Math.min(minx, step.x1); // minx = -96644
            miny = Math.min(miny, step.y1); // miny = -94584
            minz = Math.min(minz, step.z1); // minz = -96213
            maxx = Math.max(maxx, step.x2); // maxx = 95475
            maxy = Math.max(maxy, step.y2); // maxy = 93300
            maxz = Math.max(maxz, step.z2); // maxz = 97112
        }
    }

    private static int NEG = 100_1000; // 50, 100_1000
    private static int POS = 100_1000; // 50, 100_1000
    private static int SIZE = POS + NEG + 1;

    public static class Step {
        public  int x1;
        public int x2;
        public int y1;
        public int y2;
        public int z1;
        public int z2;
        public boolean on;

        public Step(int x1, int x2, int y1, int y2, int z1, int z2, boolean on) {
            this.x1 = x1;
            this.x2 = x2;
            this.y1 = y1;
            this.y2 = y2;
            this.z1 = z1;
            this.z2 = z2;
            this.on = on;
        }
    }

    private Scanner getScanner() {
        InputStream in = getClass().getResourceAsStream("input.txt");
        return new Scanner(in);
    }
}
