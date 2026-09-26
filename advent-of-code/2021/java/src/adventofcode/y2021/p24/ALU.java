package adventofcode.y2021.p24;

import java.io.InputStream;
import java.time.Duration;
import java.time.Instant;
import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Scanner;
import java.util.stream.Collectors;
import static java.lang.Integer.parseInt;
import static java.lang.Long.parseLong;
import static java.lang.StrictMath.addExact;
import static java.lang.StrictMath.floorDiv;
import static java.lang.StrictMath.floorMod;
import static java.lang.StrictMath.multiplyExact;
import static java.util.stream.Collectors.toList;

public class ALU {

    // 92967699949891, 91411143612181
    public static void main(String[] args) throws RuntimeException {
        new ALU().process();
    }

    private void process() {
        Instant one = Instant.now();
        Scanner s = getScanner();
        var ins = new ArrayList<Instruction>();
        while (s.hasNext()) {
            ins.add(new Instruction(s.nextLine()));
        }

        run();
        // var z = oldSpecific("92967699949891");
        // run3(ins);
        // old(ins);

        Instant two = Instant.now();
        Duration dur = Duration.between(one, two);
        System.out.println(dur);
    }

    private void run() {
        var zz = new long[] {1, 1, 1, 26, 1, 1, 26, 26, 1, 26, 1, 26, 26, 26};
        var xx = new long[] {11, 14, 13, -4, 11, 10, -4, -12, 10, -11, 12, -1, 0, -11};
        var yy = new long[] {3, 7, 1, 6, 14, 7, 9, 9, 6, 4, 0, 7, 12, 1};

        var states = new ArrayList<Step>();
        states.add(new Step());

        // int[] num =  {9, 2, 9, 6, 7, 6, 9, 9, 9, 4, 9, 8, 9, 1 };
        // int pos = 0;

        for (int i = 0; i < 14; i++) {
            var oldStates = dedup(states);
            states = new ArrayList<Step>();
            for (var state: oldStates) {
                //int w = num[pos++];
                for (int w = 1; w <= 9; w++) {
                    long x, y, z = state.z;
                    x = (z % 26 + xx[i]) == w ? 0 : 1;
                    z /= zz[i];
                    y = 25 * x + 1;
                    z *= y;
                    y = (w + yy[i]) * x;
                    z += y;
                    var newState = new Step();
                    newState.z = z;
                    newState.min = state.min * 10 + w;
                    newState.max = state.max * 10 + w;
                    states.add(newState);
                }
            }
        }

        var good = states.stream().filter(state -> state.z == 0).collect(toList());

        long min = Long.MAX_VALUE, max = Long.MIN_VALUE;
        for (var g: good) {
            if (min > g.min) min = g.min;
            if (max < g.max) max = g.max;
        }

        System.out.printf("\nRESULT: %d, %d\n", min, max);
    }

    private void run2(ArrayList<Instruction> ins) {
        long total = 0;
        var states = new ArrayList<Result>();
        states.add(new Result());

        for (var in: ins) {
            switch (in.command) {
                case "inp":
                    var oldStates = dedupOld(states);
                    states = new ArrayList<Result>();
                    for (var state: oldStates)
                        for (int i = 1; i <= 9; i++) {
                            var newState = new Result();
                            newState.registers[3] = state.registers[3];
                            newState.registers[in.inputPos] = i;
                            newState.min = state.min * 10 + i;
                            newState.max = state.max * 10 + i;
                            states.add(newState);
                        }
                    break;
                case "add":
                    for (var state: states) state.registers[in.inputPos] += in.valOutput(state.registers);
                    break;
                case "mul":
                    for (var state: states) state.registers[in.inputPos] *= in.valOutput(state.registers);
                    break;
                case "div":
                    for (var state: states) state.registers[in.inputPos] /= in.valOutput(state.registers);
                    break;
                case "mod":
                    for (var state: states) state.registers[in.inputPos] %= in.valOutput(state.registers);
                    break;
                case "eql":
                    for (var state: states) state.registers[in.inputPos] = state.registers[in.inputPos] == in.valOutput(state.registers) ? 1 : 0;
                    break;
                default: throw new RuntimeException("unknown: " + in.command);
            }
        }

        var good = states.stream().filter(state -> state.registers[3] == 0).collect(toList());

        long min = Long.MAX_VALUE, max = Long.MIN_VALUE;
        for (var g: good) {
            if (min > g.min) min = g.min;
            if (max < g.max) max = g.max;
        }

        System.out.printf("\nRESULT: %d, %d\n", min, max);
    }

    private void run3(ArrayList<Instruction> ins) {
        long total = 0;
        var states = new ArrayList<Result>();
        states.add(new Result());

        for (var in: ins) {
            switch (in.command) {
                case "inp":
                    var oldStates = dedupOld(states);
                    states = new ArrayList<Result>();
                    for (var state: oldStates)
                        for (int i = 1; i <= 9; i++) {
                            var newState = new Result();
                            newState.registers[3] = state.registers[3];
                            newState.registers[in.inputPos] = i;
                            newState.min = state.min * 10 + i;
                            newState.max = state.max * 10 + i;
                            states.add(newState);
                        }
                    break;
                case "add":
                    for (var state: states) state.registers[in.inputPos] = addExact(state.registers[in.inputPos], in.valOutput(state.registers));
                    break;
                case "mul":
                    for (var state: states) state.registers[in.inputPos] = multiplyExact(state.registers[in.inputPos], in.valOutput(state.registers));
                    break;
                case "div":
                    for (var state: states) state.registers[in.inputPos] = floorDiv(state.registers[in.inputPos], in.valOutput(state.registers));
                    break;
                case "mod":
                    for (var state: states) state.registers[in.inputPos] = floorMod(state.registers[in.inputPos], in.valOutput(state.registers));
                    break;
                case "eql":
                    for (var state: states) state.registers[in.inputPos] = state.registers[in.inputPos] == in.valOutput(state.registers) ? 1 : 0;
                    break;
                default: throw new RuntimeException("unknown: " + in.command);
            }
        }

        var good = states.stream().filter(state -> state.registers[3] == 0).collect(toList());

        long min = Long.MAX_VALUE, max = Long.MIN_VALUE;
        for (var g: good) {
            if (min > g.min) min = g.min;
            if (max < g.max) max = g.max;
        }

        System.out.printf("\nRESULT: %d, %d\n", min, max);
    }

    private ArrayList<Step> dedup(ArrayList<Step> states) {
        var map = new HashMap<Long, Step>();
        for (var state: states) {
            var value = map.get(state.z);
            if (value == null) {
                map.put(state.z, state);
            } else {
                value.min = Math.min(value.min, state.min);
                value.max = Math.max(value.max, state.max);
            }
        }
        var ret = new ArrayList<>(map.values());
        System.out.printf("%d:%d:%d, ", states.size(), ret.size(), states.size() - ret.size());
        return ret;
    }

    public static class Step {
        public long min, max, z;
    }

    public static class Result {
        public long min, max;
        public long[] registers = new long[4];
    }

    private ArrayList<Result> dedupOld(ArrayList<Result> states) {
        var map = new HashMap<Long, Result>();
        for (var state: states) {
            var value = map.get(state.registers);
            if (value == null) {
                map.put(state.registers[3], state);
            } else {
                value.min = Math.min(value.min, state.min);
                value.max = Math.max(value.max, state.max);
            }
        }
        var ret = new ArrayList<>(map.values());
        System.out.printf("%d:%d:%d, ", states.size(), ret.size(), states.size() - ret.size());
        return ret;
    }

    public static class Instruction {
        private final String[] split;
        public final int inputPos;
        public final String command;

        public Instruction(String line) {
            this.split = line.split(" ");
            this.command = split[0];
            this.inputPos = posVar(split[1]);
        }

        public long valOutput(long[] registers) {
            var pos = posVar(split[2]);
            return pos >= 0 ? registers[pos] : parseLong(split[2]);
        }

        public int posVar(String input) {
            switch (input) {
                case "w": return 0;
                case "x": return 1;
                case "y": return 2;
                case "z": return 3;
                default: return -1;
            }
        }
    }

    public static class State {
        private final String input;
        private int pos;
        public long[] registers = new long[4];

        public State(String input) {
            this.input = input;
        }

        public int nextInt() {
            pos++;
            return parseInt(input.substring(pos-1, pos));
        }
    }

    private void old(List<Instruction> ins) {
          var i = 92967699949891L;
          int pos = 0;
//        for (var i = 99_999_999_999_999L; ; i--) {
            var num = Long.toString(i);
            if (!num.contains("0")) {
                var state = new State(num);
                for  (var in: ins) {
                    //in.apply(state);
                    switch (in.command) {
                        case "inp":
                            state.registers[in.inputPos] = num.charAt(pos++) - '0';
                            break;
                        case "add":
                            state.registers[in.inputPos] = addExact(state.registers[in.inputPos], in.valOutput(state.registers));
                            break;
                        case "mul":
                            state.registers[in.inputPos] = multiplyExact(state.registers[in.inputPos], in.valOutput(state.registers));
                            break;
                        case "div":
                            state.registers[in.inputPos] = floorDiv(state.registers[in.inputPos], in.valOutput(state.registers));
                            break;
                        case "mod":
                            state.registers[in.inputPos] = floorMod(state.registers[in.inputPos], in.valOutput(state.registers));
                            break;
                        case "eql":
                            state.registers[in.inputPos] = state.registers[in.inputPos] == in.valOutput(state.registers) ? 1 : 0;
                            break;
                    }
                }
                boolean valid = state.registers[3] == 0;
                if (valid) {
                    System.out.println(num);
                    return;
                }
            }
        //}
    }

    private long oldSpecific(String num) {
        var zz = new long[] {1, 1, 1, 26, 1, 1, 26, 26, 1, 26, 1, 26, 26, 26};
        var xx = new long[] {11, 14, 13, -4, 11, 10, -4, -12, 10, -11, 12, -1, 0, -11};
        var yy = new long[] {3, 7, 1, 6, 14, 7, 9, 9, 6, 4, 0, 7, 12, 1};

        long w = 0, x = 0, y = 0, z = 0;
        for (var i = 0; i < 14; i++) {
            w = parseLong(num.substring(i, i+1));
            x = (z % 26 + xx[i]) == w ? 0 : 1;
            z /= zz[i];
            y = 25 * x + 1;
            z *= y;
            y = (w + yy[i]) * x;
            z += y;
        }
        return z;
    }

    private Scanner getScanner() {
        InputStream in = getClass().getResourceAsStream("input.txt");
        return new Scanner(in);
    }
}
