package adventofcode.y2021.p16;

import java.io.InputStream;
import java.time.Duration;
import java.time.Instant;
import java.util.ArrayList;
import java.util.BitSet;
import java.util.List;
import java.util.Scanner;

public class Decoder {

    public static void main(String[] args) throws RuntimeException {
        new Decoder().process();
    }

    // 940,  13476220616073
    private void process() {
        Instant one = Instant.now();
        Scanner s = getScanner();
        while (s.hasNext()) {
            // int versionSum = run(s.nextLine()).versionSum();
            var valueSum = run(s.nextLine()).valueSum();
            System.out.printf("%d\n", valueSum);
        }
        Instant two = Instant.now();
        Duration res = Duration.between(one, two);
        System.out.println(res);

    }

    private Packet run(String line) {
        return new Parser(line).nextPacket();
    }

    private final static int SIZE = 4;

    public static class Parser {
        public BitSet bits = new BitSet();
        public int total, cur;

        public Parser(String line) {
            total = line.length() * SIZE;
            for (int i=0; i<line.length(); i++) {
                var num = BitSet.valueOf(new byte[]{Byte.parseByte(line.substring(i, i+1), 16)});
                for (int j=0; j<SIZE; j++)
                    bits.set(total - 1 - (i*SIZE  + SIZE - j - 1), num.get(j));
            }
            cur = total;
        }

        public long nextLong(int nbits) {
            cur -= nbits;
            var ret= bits.get(cur, cur+nbits).toLongArray();
            return ret.length==0 ? 0 : ret[0];
        }

        public byte nextByte(int nbits) {
            return (byte) nextLong(nbits);
        }

        public Packet nextPacket() {
            int begin = cur;
            var version = nextByte(3);
            var type = nextByte(3);
            switch (type) {
                case TYPE_LITERAL:
                    byte stop = 1;
                    long number = 0;
                    while (stop == 1) {
                        stop = nextByte(1);
                        var part = nextByte(4);
                        number <<= 4;
                        number += part;
                    }
                    return new PacketNumber(version, type, begin-cur, number);
                default:
                    var list = new ArrayList<Packet>();
                    byte mode = nextByte(1);
                    if (mode == 0) {
                        var len = (int)nextLong(15);
                        int last = cur-len;
                        while (cur>last) {
                            list.add(nextPacket());
                        }
                    } else {
                        var num = nextLong(11);
                        for (int i=0; i<num; i++) {
                            list.add(nextPacket());
                        }
                    }
                    var packet = new PacketList(version, type, begin-cur, list);
                    return packet;
            }
        }
    }

    public static abstract class Packet {
        public final byte version;
        public final byte type;
        public final int size;

        public Packet(byte version, byte type, int size) {
            this.version = version;
            this.type = type;
            this.size = size;
        }

        public abstract int versionSum();
        public abstract long valueSum();
    }

    public static class PacketNumber extends Packet {
        public final long number;

        public PacketNumber(byte version, byte type, int size, long number) {
            super(version, type, size);
            this.number = number;
        }

        @Override
        public int versionSum() {
            return version;
        }

        @Override
        public long valueSum() {
            return number;
        }
    }

    public static class PacketList extends Packet {
        public List<Packet> children;

        public PacketList(byte version, byte type, int size, List<Packet> children) {
            super(version, type, size);
            this.children = children;
        }

        @Override
        public int versionSum() {
            return children.stream().mapToInt(Packet::versionSum).sum() + version;
        }

        @Override
        public long valueSum() {
            switch (type) {
                case TYPE_SUM:
                    return children.stream().mapToLong(Packet::valueSum).sum();
                case TYPE_PRODUCT:
                    return children.stream().mapToLong(Packet::valueSum).reduce(1, (a,b)->a*b);
                case TYPE_MIN:
                    return children.stream().mapToLong(Packet::valueSum).min().getAsLong();
                case TYPE_MAX:
                    return children.stream().mapToLong(Packet::valueSum).max().getAsLong();
                case TYPE_GREATER:
                    assert children.size() == 2;
                    return children.get(0).valueSum() > children.get(1).valueSum() ? 1 : 0;
                case TYPE_LESS:
                    assert children.size() == 2;
                    return children.get(0).valueSum() < children.get(1).valueSum() ? 1 : 0;
                case TYPE_EQUAL:
                    assert children.size() == 2;
                    return children.get(0).valueSum() == children.get(1).valueSum() ? 1 : 0;
                default:
                    throw new RuntimeException("unknown type");
            }
        }
    }

    public static final byte TYPE_SUM = 0;
    public static final byte TYPE_PRODUCT = 1;
    public static final byte TYPE_MIN = 2;
    public static final byte TYPE_MAX = 3;
    public static final byte TYPE_LITERAL = 4;
    public static final byte TYPE_GREATER = 5;
    public static final byte TYPE_LESS = 6;
    public static final byte TYPE_EQUAL = 7;

    private Scanner getScanner() {
        InputStream in = getClass().getResourceAsStream("input.txt");
        return new Scanner(in);
    }
}
