package adventofcode.y2021.p21;

import java.util.ArrayDeque;
import java.util.HashMap;
import java.util.Stack;

public class Dice {

    public static void main(String[] args) throws RuntimeException {
        new Dice().process();
    }

    private void process() {
        // var a = run1(4, 8);
        // var a = run1(5, 8);
        //System.out.println(a);

        //run2(4, 8);
        run2(5, 8);
    }

    // 1067724
    private int run1(int pos1, int pos2) {
        int score1 = 0, score2 = 0;
        while (true) {
            pos1 = (pos1 + dice() + dice() + dice() - 1) % 10 + 1;
            score1 += pos1;
            if (score1 >= 1000) {
                return times * score2;
            }

            pos2 = (pos2 + dice() + dice() + dice() - 1) % 10 + 1;
            score2 += pos2;
            if (score2 >= 1000) {
                return times * score1;
            }
        }
    }

    private int dice() {
        ++times;
        return (++_dice - 1) % 100 + 1;
    }

    private int times = 0;
    private int _dice = 0;

    // 630947104784464
    private void run2(int pos1, int pos2) {
        int[] diceNumber = { 3, 4, 5, 6, 7, 8, 9 };
        int[] diceTimes = { 1, 3, 6, 7, 6, 3, 1 };

        long total1 = 0, total2 = 0;
        var list = new ArrayDeque<State>();
        list.add(new State(pos1, pos2, 0, 0, 1,true));
        while (!list.isEmpty()) {
            var cur = list.remove();
                for (int i = 0; i<diceNumber.length; i++) {
                    long newTimes = cur.times * diceTimes[i];
                    if (cur.is1) {
                        int newPos1 = (cur.pos1 + diceNumber[i] - 1) % 10 + 1;
                        int newScore1 = cur.score1 + newPos1;
                        if (newScore1 >= 21) {
                            total1 += newTimes;
                        } else {
                            list.add(new State(newPos1, cur.pos2, newScore1, cur.score2, newTimes, !cur.is1));
                        }
                    } else {
                        int newPos2 = (cur.pos2 + diceNumber[i] - 1) % 10 + 1;
                        int newScore2 = cur.score2 + newPos2;
                        if (newScore2 >= 21) {
                            total2 += newTimes;
                        } else {
                            list.add(new State(cur.pos1, newPos2, cur.score1, newScore2, newTimes, !cur.is1));
                        }
                    }
                }
        }
        // RES1: 1521956905244, RES2: 1537725478079, MAX: 1537725478079
        System.out.printf("RES1: %d, RES2: %d, MAX: %d", total1, total2, Math.max(total1, total2));
    }

    public static class State {
        int pos1, pos2, score1, score2;
        long times;
        boolean is1;

        public State(int pos1, int pos2, int score1, int score2, long times, boolean is1) {
            this.pos1 = pos1;
            this.pos2 = pos2;
            this.score1 = score1;
            this.score2 = score2;
            this.times = times;
            this.is1 = is1;
        }
    }

}
