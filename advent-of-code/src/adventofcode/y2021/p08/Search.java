package adventofcode.y2021.p08;

import java.io.InputStream;
import java.util.ArrayList;
import java.util.Arrays;
import java.util.Collections;
import java.util.HashSet;
import java.util.LinkedList;
import java.util.List;
import java.util.Scanner;
import java.util.Set;
import java.util.stream.Collectors;
import java.util.stream.IntStream;
import java.util.stream.Stream;
import static java.util.stream.Collectors.toList;

public class Search {

    public static void main(String[] args) throws RuntimeException {
        new Search().process();
    }

    private void process() {
        Scanner s = getScanner();
        while (s.hasNext()) {
            var in = new String[10];
            for (int i=0; i<in.length; i++) {
                in[i] =s.next();
            }
            s.next();
            var out = new String[4];
            for (int i=0; i<out.length; i++) {
                out[i] = s.next();
            }
            s.nextLine();
            run2(in, out);
        }
        System.out.printf("%d ", total);
    }

    private int total;

    // 284
    private void run1(String[] ins, String[] outs) {
        Permutations.of(Arrays.asList(0, 1, 2, 3, 4, 5, 6)).forEach(p -> {
            var map = p.mapToInt(Integer::intValue).toArray();
            if (Arrays.stream(ins).allMatch(in -> SET.contains(transform(map, in)))) {
                for (var out : outs) {
                    if (HITS.contains(transform(map, out))) {
                        total++;
                    }
                }
            }
        });
    }

    // 973499
    private void run2(String[] ins, String[] outs) {
        Permutations.of(Arrays.asList(0, 1, 2, 3, 4, 5, 6)).forEach(p -> {
            var map = p.mapToInt(Integer::intValue).toArray();
            if (Arrays.stream(ins).allMatch(in -> SET.contains(transform(map, in)))) {
                int num = 0;
                for (var out : outs) {
                    num = num * 10 + LIST.indexOf(transform(map, out));
                }
                total += num;
            }
        });
    }

    private String transform(int[] map, String in) {
        char[] check = in.toCharArray();
        for (int i = 0; i < check.length; i++) {
            check[i] = (char) (map[check[i] - 'a'] + 'a');
        }
        Arrays.sort(check);
        return new String(check);
    }

    private final static String[] NUMBERS =
            { "abcefg", "cf", "acdeg", "acdfg", "bcdf", "abdfg", "abdefg", "acf", "abcdefg", "abcdfg"};

    private final static List<String> LIST = Arrays.asList(NUMBERS);
    private final static Set<String> SET = new HashSet<>(LIST);
    private final static Set<String> HITS = new HashSet<>(Arrays.asList(NUMBERS[1], NUMBERS[4], NUMBERS[7], NUMBERS[8]));


    public static class Permutations {

        public static <T> Stream<Stream<T>> of(final List<T> items) {
            return IntStream.range(0, factorial(items.size())).mapToObj(i -> permutation(i, items).stream());
        }

        private static int factorial(final int num) {
            return IntStream.rangeClosed(2, num).reduce(1, (x, y) -> x * y);
        }

        private static <T> List<T> permutation(final int count, final LinkedList<T> input, final List<T> output) {
            if (input.isEmpty()) { return output; }

            final int factorial = factorial(input.size() - 1);
            output.add(input.remove(count / factorial));
            return permutation(count % factorial, input, output);
        }

        private static <T> List<T> permutation(final int count, final List<T> items) {
            return permutation(count, new LinkedList<>(items), new ArrayList<>());
        }

    }

    private Scanner getScanner() {
        InputStream in = getClass().getResourceAsStream("input.txt");
        return new Scanner(in);
    }
}
