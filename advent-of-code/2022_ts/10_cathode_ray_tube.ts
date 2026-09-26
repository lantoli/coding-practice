import { readFileSync } from 'fs';

const inputContent = readFileSync("inputs/10_input.txt", 'utf-8');
//const inputContent = readFileSync("inputs/10_sample.txt", 'utf-8');
const lines = inputContent.split('\n');
lines.pop(); // Remove last empty line

const wide = 40, high = 6;
const when = [20, 60, 100, 140, 180, 220];
const values: number[] = [];
const crt: boolean[] = new Array(wide * high);
let x = 1, cycle = 0;

const update = () => { update1(); update2(); };

function update1() { if (cycle == when[values.length]) values.push(x) };
function result1() { console.log(when.map((w, idx) => w * values[idx]).reduce((a, b) => a + b)) };

function update2() {
    const pos = cycle - 1, posLine = pos % wide;
    crt[pos] = Math.abs(posLine - x) < 2;
};
function result2() {
    crt.forEach((val, idx) => {
        if (idx > 0 && idx % wide == 0) process.stdout.write("\n");
        process.stdout.write(val ? "#" : ".");
    });
    process.stdout.write("\n")
}

function run() {
    for (var line of lines) {
        const [instr, arg] = line.split(" ");
        if (instr === "noop") {
            cycle++; update();
        } else if (instr === "addx") {
            cycle++; update(); cycle++; update();
            x += Number(arg);
        } else throw new Error("unexpected");
    }
    result1();
    result2();
};

run(); // 15360, PHLHJGZA
/*
###..#..#.#....#..#...##..##..####..##..
#..#.#..#.#....#..#....#.#..#....#.#..#.
#..#.####.#....####....#.#......#..#..#.
###..#..#.#....#..#....#.#.##..#...####.
#....#..#.#....#..#.#..#.#..#.#....#..#.
#....#..#.####.#..#..##...###.####.#..#.
*/
