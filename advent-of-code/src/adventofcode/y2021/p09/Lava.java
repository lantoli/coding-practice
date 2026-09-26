package adventofcode.y2021.p09;

import java.io.InputStream;
import java.util.ArrayList;
import java.util.Arrays;
import java.util.Comparator;
import java.util.HashSet;
import java.util.List;
import java.util.Scanner;
import java.util.Set;

public class Lava {

    public static void main(String[] args) throws RuntimeException {
        new Lava().process();
    }

    private void process() {
        int[][] grid = new int[ROWS][COLS];
        Scanner s = getScanner();
        for (int i=0; i<ROWS; i++) {
            String line = s.nextLine();

            for (int j = 0; j < COLS; j++)
                grid[i][j] = line.charAt(j) - '0';
        }
        run2(grid);
    }

    // 591
    private void run1(int[][] grid) {
        int risk = 0;
        for (int i=0; i<ROWS; i++) {
            for (int j = 0; j < COLS; j++) {
                if (isLow(grid, i, j)) {
                    risk += 1 + grid[i][j];
                }
            }
        }
        System.out.printf("risk: %d ", risk);
    }

    //  1113424: 106, 104, 101
    private void run2(int[][] grid) {
        var basins = new ArrayList<Integer>();
        for (int i=0; i<ROWS; i++) {
            for (int j = 0; j < COLS; j++) {
                if (isLow(grid, i, j)) {
                    int s = size(grid, new HashSet<>(), i, j);
                    basins.add(s);
                    System.out.printf("%d, ", s);
                }
            }
        }
        basins.sort(Comparator.reverseOrder());
        System.out.println("\nOrdered:");
        System.out.println(Arrays.toString(basins.toArray()));
    }

    private int size(int[][] grid, Set<Integer> visited, int i, int j) {
        if (i<0 || i>ROWS-1 || j<0 || j>COLS-1 || grid[i][j] == 9 || visited.contains(i*COLS + j)) return 0;
        visited.add(i*COLS + j);
        return 1 + size(grid,visited,  i-1, j) + size(grid, visited, i+1, j) + size(grid, visited, i, j-1) + size(grid, visited, i, j+1);
    }

    private boolean isLow(int[][] grid, int i, int j) {
        if (i>0 && grid[i-1][j] <= grid[i][j]) return false;
        if (i<ROWS-1 && grid[i+1][j] <= grid[i][j]) return false;
        if (j>0 && grid[i][j-1] <= grid[i][j]) return false;
        if (j<COLS-1 && grid[i][j+1] <= grid[i][j]) return false;
        return true;
    }

    private static final int COLS = 100; // 10;
    private static final int ROWS = 100; // 5;

    private Scanner getScanner() {
        InputStream in = getClass().getResourceAsStream("input.txt");
        return new Scanner(in);
    }
}
