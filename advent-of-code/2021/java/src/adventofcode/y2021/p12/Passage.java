package adventofcode.y2021.p12;

import java.io.InputStream;
import java.util.ArrayDeque;
import java.util.ArrayList;
import java.util.HashMap;
import java.util.HashSet;
import java.util.List;
import java.util.Map;
import java.util.Queue;
import java.util.Scanner;
import java.util.Set;
import static java.lang.Character.isUpperCase;

public class Passage {

    public static void main(String[] args) throws RuntimeException {
        new Passage().process();
    }

    private void process() {
        Scanner s = getScanner();
        var edges = new ArrayList<String[]>();
        while (s.hasNext()) {
            String[] edge = s.nextLine().split("-");
            edges.add(edge);
            edges.add(new String[] { edge[1], edge[0] });
        }
        run(edges);
    }

    // 4773, alreadySmall (part 2): 116985
    private void run(List<String[]> edges) {
        edges.removeIf(edge -> edge[1].equals("start"));
        int total = 0;
        Queue<State> q = new ArrayDeque<>();
        q.add(new State());
        while (!q.isEmpty()) {
            var current = q.remove();
            for (String[] edge: edges) {
                if (edge[0].equals(current.last)) {
                    String dest = edge[1];
                    boolean big = isUpperCase(dest.charAt(0));
                    if (dest.equals("end")) {
                        total++;
                    } else if (big || !current.alreadySmall || !current.visitedSmall.contains(dest)) {
                        q.add(new State(current, dest, big));
                    }
                }
            }
        }
        System.out.printf("total: %d\n", total);
    }

    public static class State {
        public final String last;
        public final Set<String> visitedSmall = new HashSet<>();
        public final List<String> list = new ArrayList<>();
        public boolean alreadySmall;

        public State() {
            last = "start";
        }

        public State(State current, String dest, boolean big) {
            last = dest;
            list.add(last);
            visitedSmall.addAll(current.visitedSmall);
            alreadySmall = current.alreadySmall || visitedSmall.contains(dest);
            if (!big) visitedSmall.add(dest);
        }
    }


    private Scanner getScanner() {
        InputStream in = getClass().getResourceAsStream("input.txt");
        return new Scanner(in);
    }
}
