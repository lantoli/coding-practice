package adventofcode.y2021.p13;

import java.io.InputStream;
import java.util.Scanner;
import static java.lang.Integer.parseInt;

public class Origami {

    public static void main(String[] args) throws RuntimeException {
        new Origami().process();
    }

    // 731, ZKAUCFUC
    private void process() {
        Scanner s = getScanner();
        var grid = new boolean[SIZE][SIZE];
        String line;
        while (!(line = s.nextLine()).isEmpty()) {
            var strPoint = line.split(",");
            grid[parseInt(strPoint[1])][parseInt(strPoint[0])] = true;
        }
        while (s.hasNext()) {
            var fold = s.nextLine().split(" ")[2].split("=");
            boolean foldY = fold[0].equals("y");
            int where = parseInt(fold[1]);
            run(grid, foldY, where);
            String a = "";
        }
        // run(grid, true, 7);
        // run(grid, false, 5);
        // run(grid, false, 655);

        int total = 0;
        for (int y=0; y<6; y++) {
            for (int x = 0; x < 40; x++)
                System.out.print(grid[y][x] ? '#' : ' ');
            System.out.println();
        }
        System.out.printf("total: %d\n", total);
    }

    private void run(boolean[][] grid, boolean foldY, int line) {

        for (int y=0; y<line; y++)
            for (int x=0; x<SIZE; x++)
                if (foldY)
                    grid[y][x] |= grid[2*line-y][x];
                else
                    grid[x][y] |= grid[x][2*line-y];
    }

    private static int SIZE = 1311; // 15

    private Scanner getScanner() {
        InputStream in = getClass().getResourceAsStream("input.txt");
        return new Scanner(in);
    }
}
