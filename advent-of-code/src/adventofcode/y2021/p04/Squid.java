package adventofcode.y2021.p04;

import java.io.InputStream;
import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;
import java.util.Scanner;

public class Squid {

    public static void main(String[] args) throws RuntimeException {
        new Squid().process();
    }

    private void process() {
        Scanner s = getScanner();
        int[] numbers = Arrays.stream(s.nextLine().split(",")).mapToInt(Integer::parseInt).toArray();
        List<Board> boards = new ArrayList();
        while (s.hasNext()) {
            int[][] grid = new int[SIZE][SIZE];
            for (int j=0; j<SIZE; j++) {
                for (int i=0; i<SIZE; i++) {
                    grid[j][i] = s.nextInt();
                }
            }
            boards.add(new Board(grid));
        }
        run2(numbers, boards);
    }

    // 44088
    private void run1(int[] numbers, List<Board> boards) {

        for (int number: numbers) {
            for (var board: boards) {
                if (board.apply(number)) {
                    System.out.printf("total: %d\n", number * board.sum());
                    return;
                }
            }
        }
    }

    // 23670
    private void run2(int[] numbers, List<Board> boards) {

        for (int number: numbers) {
            Board last = null;
            var it = boards.iterator();
            while (it.hasNext()) {
                Board board = it.next();
                if (board.apply(number)) {
                    last = board;
                    it.remove();
                }
            }
            if (boards.size() == 0) {
                System.out.printf("total: %d\n", number * last.sum());
                return;
            }
        }
    }

    public static class Board {
        public final int[][] grid;
        public boolean[][] hit = new boolean[SIZE][SIZE];
        public int[] cols = new int[SIZE];
        public int[] rows = new int[SIZE];
        public Board(int[][] grid) {
            this.grid = grid;
        }

        public boolean apply(int number) {
            for (int j=0; j<SIZE; j++) {
                for (int i=0; i<SIZE; i++) {
                    if (grid[j][i] == number) {
                        hit[j][i] = true;
                        if (++rows[j] == SIZE) return true;
                        if (++cols[i] == SIZE) return true;
                        return false;
                    }
                }
            }
            return false;
        }

        public int sum() {
            int ret = 0;
            for (int j=0; j<SIZE; j++) {
                for (int i=0; i<SIZE; i++) {
                    if (!hit[j][i]) {
                        ret += grid[j][i];
                    }
                }
            }
            return ret;
        }
    }

    final private static int SIZE = 5;
    private Scanner getScanner() {
        InputStream in = getClass().getResourceAsStream("input.txt");
        return new Scanner(in);
    }
}
