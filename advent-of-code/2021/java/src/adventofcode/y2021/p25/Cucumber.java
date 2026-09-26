package adventofcode.y2021.p25;

import java.io.InputStream;
import java.time.Duration;
import java.time.Instant;
import java.util.ArrayDeque;
import java.util.Scanner;
import static java.lang.Integer.parseInt;
import static java.lang.Long.parseLong;

public class Cucumber {

    public static void main(String[] args) throws RuntimeException {
        new Cucumber().process();
    }

    private void process() {
        Instant one = Instant.now();
        Scanner s = getScanner();
        var grid = new char[ROWS][];
        for (int i = 0; i < ROWS; i++) {
            grid[i] = s.nextLine().toCharArray();
        }
        run(grid);

        Instant two = Instant.now();
        Duration dur = Duration.between(one, two);
        System.out.println(dur);
    }

    // 471
    private void run(char[][] grid) {
        int steps = 0;
        boolean changed = true;
        while (changed) {
            var changes1 = new ArrayDeque<int[]>();
            for (int i=0; i<ROWS; i++)
                for (int j=0; j<COLS; j++)
                    if (grid[i][j] == '>' && grid[i][(j + 1) % COLS] == '.')
                        changes1.add(new int[] { i, j, i,  (j + 1) % COLS });
            for (var change: changes1) {
                grid[change[2]][change[3]] = grid[change[0]][change[1]];
                grid[change[0]][change[1]] = '.';
            }
            var changes2 = new ArrayDeque<int[]>();
            for (int i=0; i<ROWS; i++)
                for (int j=0; j<COLS; j++)
                    if (grid[i][j] == 'v' && grid[(i + 1) % ROWS][j] == '.')
                        changes2.add(new int[] { i, j, (i + 1) % ROWS, j });
            for (var change: changes2) {
                grid[change[2]][change[3]] = grid[change[0]][change[1]];
                grid[change[0]][change[1]] = '.';
            }
            changed = !changes1.isEmpty() || !changes2.isEmpty();
            steps++;
        }
        System.out.printf("RET: %d\n", steps);
    }

    private static int ROWS = 137;  // 9, 137
    private static int COLS = 139; // 10, 139

    private Scanner getScanner() {
        InputStream in = getClass().getResourceAsStream("input.txt");
        return new Scanner(in);
    }
}
