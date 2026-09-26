package adventofcode.y2021.p15;

import java.io.InputStream;
import java.time.Duration;
import java.time.Instant;
import java.util.ArrayDeque;
import java.util.ArrayList;
import java.util.HashMap;
import java.util.HashSet;
import java.util.List;
import java.util.Scanner;
import java.util.Set;
import static java.lang.Integer.parseInt;

public class Chiton {

    public static void main(String[] args) throws RuntimeException {
        new Chiton().process();
    }

    // 581, 2916
    private void process() {
        Instant one = Instant.now();
        Scanner s = getScanner();
        int[][] grid = new int[SIZE_INPUT][SIZE_INPUT];
        for (int j=0; j<SIZE_INPUT; j++) {
            var line = s.nextLine();
            for (int i=0; i<SIZE_INPUT; i++) {
                grid[j][i] = parseInt(line.substring(i, i+1));
            }
        }
        grid = make5(grid);
        run(grid);
        Instant two = Instant.now();
        Duration res = Duration.between(one, two);
        System.out.println(res);
    }

    private int[][] make5(int[][] grid) {
        var ret = new int[SIZE][SIZE];
        for (int yy=0; yy<TIMES; yy++)
            for (int xx=0; xx<TIMES; xx++) {
                for (int y=0; y<SIZE_INPUT; y++)
                    for (int x=0; x<SIZE_INPUT; x++) {
                        int val = grid[y][x] + xx + yy;
                        if (val > 9) val -= 9;
                        ret[yy*SIZE_INPUT+y][xx*SIZE_INPUT+x] = val;
                    }
            }
        return ret;
    }

    private void run(int[][] grid) {
        int level = Integer.MAX_VALUE;
        int posEnd = pos(SIZE-1, SIZE-1);
        var states = new ArrayDeque<State>();
        states.push(new State());
        var best = new HashMap<Integer, Integer>();
        while (!states.isEmpty()) {
            var state = states.pop();
            for (int pos: state.neighbors()) {
                int newLevel = state.level + grid[y(pos)][x(pos)];
                if (pos == posEnd) {
                    level = Math.min(level, newLevel);
                } else {
                    if (newLevel >= level) continue;
                    var bestPos = best.getOrDefault(pos, Integer.MAX_VALUE);
                    if (newLevel < bestPos) {
                        best.put(pos, newLevel);
                        states.add(state.newState(pos, newLevel));
                    }
                }
            }
        }
        System.out.printf("Level: %d\n", level);
    }

    public static int pos(int y, int x) {
        return y*SIZE + x;
    }

    public static int y(int pos) {
        return pos / SIZE;
    }

    public static int x(int pos) {
        return pos % SIZE;
    }

    public static class State {
        public int level;
        public int last;

        public State newState(int pos, int newLevel) {
            var that = new State();
            that.level = newLevel;
            that.last = pos;
            return that;
        }

        public List<Integer> neighbors() {
            var ret = new ArrayList<Integer>();
            int x = x(last), y = y(last);
            if (y-1 >= 0) {
                ret.add(pos(y-1, x));
            }
            if (y+1 < SIZE) {
                ret.add(pos(y+1, x));
            }
            if (x-1 >= 0) {
                ret.add(pos(y, x-1));
            }
            if (x+1 < SIZE) {
                ret.add(pos(y, x+1));
            }
            return ret;
        }
    }

    final private static int SIZE_INPUT = 100; // 10, 100
    final private static int TIMES = 5; // 1, 5
    final private static int SIZE = SIZE_INPUT * TIMES;

    private Scanner getScanner() {
        InputStream in = getClass().getResourceAsStream("input.txt");
        return new Scanner(in);
    }
}
