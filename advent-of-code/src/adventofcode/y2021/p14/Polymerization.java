package adventofcode.y2021.p14;

import java.io.InputStream;
import java.util.HashMap;
import java.util.Scanner;
import java.util.TreeMap;
import static java.lang.Integer.parseInt;

public class Polymerization {

    public static void main(String[] args) throws RuntimeException {
        new Polymerization().process();
    }

    // 2745,   PB
    private void process() {
        Scanner s = getScanner();
        String polymer = s.nextLine();
        s.nextLine();
        var map = new HashMap<String,String>();

        while(s.hasNext()) {
            var line = s.nextLine().split(" -> ");
            map.put(line[0], line[1]);
        }
        for (int i=0; i<40; i++) {
            var count = new TreeMap<Character, Long>();
            for (char ch: polymer.toCharArray()) count.merge(ch, 1L, Long::sum);
            var min = count.values().stream().mapToLong(v -> v).min().getAsLong();
            var max = count.values().stream().mapToLong(v -> v).max().getAsLong();
            System.out.printf("Len: %d, Min: %d, Max: %d, Diff: %d", polymer.length(), min, max, max - min);

            System.out.print(", min: ");
            for (var entry: count.entrySet()) {
                if (entry.getValue() == min) {
                    System.out.print(entry.getKey());
                }
            }
            System.out.print(", max: ");
            for (var entry: count.entrySet()) {
                if (entry.getValue() == max) {
                    System.out.print(entry.getKey());
                }
            }


            System.out.printf(", pol: %s\n", polymer.substring(0, Math.min(polymer.length(), 100)));
            //System.out.printf("%d ", i);
            polymer = run(polymer, map);
        }
    }

    private String run(String polymer, HashMap<String, String> map) {
        var ret = new StringBuilder();
        ret.append(polymer.charAt(0));
        for (int i=1; i<polymer.length(); i++) {
            var add = map.get(polymer.substring(i-1, i+1));
            if (add != null) ret.append(add);
            ret.append(polymer.charAt(i));
        }
        return ret.toString();
    }


    private Scanner getScanner() {
        InputStream in = getClass().getResourceAsStream("input.txt");
        return new Scanner(in);
    }
}
