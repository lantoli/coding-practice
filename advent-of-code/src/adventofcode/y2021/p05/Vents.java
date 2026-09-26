package adventofcode.y2021.p05;

import java.io.InputStream;
import java.util.ArrayList;
import java.util.Scanner;

public class Vents {

    public static void main(String[] args) throws RuntimeException {
        new Vents().process();
    }

    private void process() {
        Scanner s = getScanner();
        var list = new ArrayList<Line>();
        while (s.hasNext()) {
            String[] p1 = s.next().split(",");
            s.next();
            String[] p2 = s.next().split(",");
            s.nextLine();
            list.add(new Line(Integer.parseInt(p1[0]), Integer.parseInt(p1[1]), Integer.parseInt(p2[0]), Integer.parseInt(p2[1])));
        }
        run2(list);
    }

    // 5774
    private void run1(ArrayList<Line> lines) {
        int[][] board = new int[SIZE][SIZE];
        for (var line: lines) {
            line.fill1(board);
        }
        int total = 0;
        for (int j = 0; j < SIZE; j++)
            for (int i = 0; i < SIZE; i++)
                if (board[j][i] > 1) total++;
        System.out.printf("total: %d\n", total);
    }

    // 18423
    private void run2(ArrayList<Line> lines) {
        int[][] board = new int[SIZE][SIZE];
        for (var line: lines) {
            line.fill2(board);
        }
        int total = 0;
        for (int j = 0; j < SIZE; j++)
            for (int i = 0; i < SIZE; i++)
                if (board[j][i] > 1) total++;
        System.out.printf("total: %d\n", total);
    }


    public static class Line {
        private final int x1;
        private final int y1;
        private final int x2;
        private final int y2;

        public Line(int x1, int y1, int x2, int y2) {
            this.x1 = x1;
            this.y1 = y1;
            this.x2 = x2;
            this.y2 = y2;
        }

        public void fill1(int[][] board) {
            int xinc = 0, yinc = 0;
            if (x1 == x2) {
                yinc = y1 < y2 ? 1 : -1;
            } else if (y1 == y2){
                xinc = x1 < x2 ? 1 : -1;
            } else {
                return;
            }
            int x = x1, y = y1;
            try {
                while (x != x2 + xinc || y != y2 + yinc) {
                    board[y][x]++;
                    x += xinc;
                    y += yinc;
                }
            } catch (ArrayIndexOutOfBoundsException a) {

            }
        }

        public void fill2(int[][] board) {
            int xinc = 0, yinc = 0;
            if (x1 < x2) {
                xinc = 1;
            } else if (x1 > x2) {
                xinc = -1;
            }

            if (y1 < y2) {
                yinc = 1;
            } else if (y1 > y2) {
                yinc = -1;
            }

            int x = x1, y = y1;
            try {
                while (x != x2 + xinc || y != y2 + yinc) {
                    board[y][x]++;
                    x += xinc;
                    y += yinc;
                }
            } catch (ArrayIndexOutOfBoundsException a) {

            }
        }
    }

    private final static int SIZE = 1000;

    private Scanner getScanner() {
        InputStream in = getClass().getResourceAsStream("input.txt");
        return new Scanner(in);
    }
}
