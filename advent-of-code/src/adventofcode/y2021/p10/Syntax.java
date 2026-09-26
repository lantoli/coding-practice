package adventofcode.y2021.p10;

import java.io.InputStream;
import java.lang.reflect.Array;
import java.util.ArrayList;
import java.util.Arrays;
import java.util.Collections;
import java.util.Comparator;
import java.util.HashSet;
import java.util.List;
import java.util.Scanner;
import java.util.Set;
import java.util.Stack;

public class Syntax {

    public static void main(String[] args) throws RuntimeException {
        new Syntax().process();
    }

    private void process() {
        Scanner s = getScanner();
        var list = new ArrayList<Long>();
        while (s.hasNext()) {
            long ret = run2(s.nextLine());
            if (ret != 0) {
                list.add(ret);
                System.out.printf("%d, ", ret);
            }
        }
        Collections.sort(list);
        System.out.printf("middle: %d, ", list.get(list.size()/2));
    }

    // 387363
    private long run1(String line) {
        var stack = new Stack<Integer>();
        for (char ch: line.toCharArray()) {
            int open = OPEN.indexOf(ch);
            int closed = CLOSED.indexOf(ch);
            if (open >= 0) {
                stack.push(open);
            } else {
                int candidate = stack.pop();
                if (candidate != closed) {
                    return VALUE[closed];
                }
            }
        }
        return 0;
    }

    // 4330777059
    private long run2(String line) {
        var stack = new Stack<Integer>();
        for (char ch: line.toCharArray()) {
            int open = OPEN.indexOf(ch);
            int closed = CLOSED.indexOf(ch);
            if (open >= 0) {
                stack.push(open);
            } else {
                int candidate = stack.pop();
                if (candidate != closed) {
                    return 0;
                }
            }
        }
        long total = 0;
        while (!stack.isEmpty()) {
            total = total * 5 + stack.pop() + 1;
        }
        return total;
    }

    private final static List<Character> OPEN = Arrays.asList('(', '[', '{', '<');
    private final static List<Character> CLOSED = Arrays.asList(')', ']', '}', '>');
    private final static int[] VALUE = {3, 57, 1197, 25137};

    private Scanner getScanner() {
        InputStream in = getClass().getResourceAsStream("input.txt");
        return new Scanner(in);
    }
}
