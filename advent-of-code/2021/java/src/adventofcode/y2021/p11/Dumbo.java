package adventofcode.y2021.p11;

import java.io.InputStream;
import java.util.HashSet;
import java.util.Scanner;
import java.util.Set;

public class Dumbo {

    public static void main(String[] args) throws RuntimeException {
        new Dumbo().process();
    }

    private void process() {
        Scanner s = getScanner();
        var grid = new int[SIZE][SIZE];
        for (int i = 0; i < SIZE; i++) {
            String line = s.nextLine();
            for (int j = 0; j < SIZE; j++) {
                grid[i][j] = line.charAt(j) - '0';
            }
        }
        run2(grid);
    }

    // 1642
    private void run1(int[][] grid) {
        int flashes = 0;
        for (int s = 0; s < STEPS; s++) {
            flashes += step(grid);
        }
        System.out.printf("flashes: %d\n", flashes);
    }

    // 320
    private void run2(int[][] grid) {
        for (int step = 1; ; step++) {
            if(step(grid) == 100) {
                System.out.printf("step: %d\n", step);
                return;
            }
        }
    }

    private int step(int[][] grid) {
        var visited = new HashSet<Integer>();
        for (int i = 0; i < SIZE; i++) {
            for (int j = 0; j < SIZE; j++) {
                ++grid[i][j];
                propagate(grid, visited, i, j);
            }
        }
        int flashes = 0;
        for (int i = 0; i < SIZE; i++)
            for (int j = 0; j < SIZE; j++)
                if(grid[i][j] > 9) {
                    grid[i][j] = 0;
                    flashes++;
                }
        return flashes;
    }

    private int propagate(int[][] grid, Set<Integer> visited, int i, int j) {
        if (grid[i][j] <= 9 || visited.contains(i*SIZE+j)) return 0;
        visited.add(i*SIZE+j);
        int total = 1;
        for (int ii = i-1; ii <= i+1; ii++)
            for (int jj = j-1; jj <= j+1; jj++)
                if (ii >= 0 && ii < SIZE && jj >=0 && jj < SIZE && grid[ii][jj] <= 9) {
                    if (++grid[ii][jj] > 9) {
                        total += propagate(grid, visited, ii, jj);
                    }
                }
        return total;
    }

    private final static int SIZE = 10;
    private final static int STEPS = 100;

    private Scanner getScanner() {
        InputStream in = getClass().getResourceAsStream("input.txt");
        return new Scanner(in);
    }
}
