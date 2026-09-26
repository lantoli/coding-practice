// NOTE: input is hardcoded in code so it wont work with a different input

type Monkey = {
    items: number[],
    module: number,
    if: number,
    else: number,
    op: (old: bigint) => bigint,
    count: number,
}

const monkeysSample: Monkey[] = [
    {items: [79, 98], module: 23, if: 2, else: 3, op: (old: bigint) => old * 19n, count: 0 },
    {items: [54, 65, 75, 74], module: 19, if: 2, else: 0, op: (old: bigint) => old + 6n, count: 0 },
    {items: [79, 60, 97], module: 13, if: 1, else: 3, op: (old: bigint) => old * old, count: 0 },
    {items: [74], module: 17, if: 0, else: 1, op: (old: bigint) => old + 3n, count: 0 },
];

const monkeysInput: Monkey[] = [
    {items: [59, 74, 65, 86], module: 7, if: 6, else: 2, op: (old: bigint) => old * 19n, count: 0 },
    {items: [62, 84, 72, 91, 68, 78, 51], module: 2, if: 2, else: 0, op: (old: bigint) => old + 1n, count: 0 },
    {items: [78, 84, 96], module: 19, if: 6, else: 5, op: (old: bigint) => old + 8n, count: 0 },
    {items: [97, 86], module: 3, if: 1, else: 0, op: (old: bigint) => old * old, count: 0 },
    {items: [50], module: 13, if: 3, else: 1, op: (old: bigint) => old + 6n, count: 0 },
    {items: [73, 65, 69, 65, 51], module: 11, if: 4, else: 7, op: (old: bigint) => old * 17n, count: 0 },
    {items: [69, 82, 97, 93, 82, 84, 58, 63], module: 5, if: 5, else: 7, op: (old: bigint) => old + 5n, count: 0 },
    {items: [81, 78, 82, 76, 79, 80], module: 17, if: 3, else: 4, op: (old: bigint) => old + 3n, count: 0 },
];

function run(rounds: number, relief: boolean, monkeys: Monkey[]) {
    monkeys = monkeys.map(monkey => ({...monkey, items: [...monkey.items]})); // deep copy
    const mod = BigInt(monkeys.map(monkey => monkey.module).reduce((a, b) => a * b));
    for (var round = 0; round < rounds; round++) {
        for (let monkey of monkeys) {
            monkey.count += monkey.items.length;
            for (let item of monkey.items) {
                let worried = monkey.op(BigInt(item));
                if (relief) worried /= 3n;
                const num = Number(worried % mod);
                const dest = num % monkey.module == 0 ? monkey.if : monkey.else;
                monkeys[dest].items.push(num);
            };
            monkey.items.length = 0;
        };
    }
    const counts = monkeys.map(monkey => monkey.count).sort((a, b) => a - b);
    console.debug(counts.pop()! * counts.pop()!);
};

run(20, true, monkeysSample); // 10605 (sample)
run(10_000, false, monkeysSample); // 2713310158 (sample)
run(20, true, monkeysInput); // 61005
run(10_000, false, monkeysInput); // 20567144694
