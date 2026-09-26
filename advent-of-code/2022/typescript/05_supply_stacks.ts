import { readFileSync } from 'fs';

const inputContent = readFileSync("inputs/05_input.txt", 'utf-8');
const lines = inputContent.split('\n');
lines.pop(); // Remove last empty line

type Move = { times: number, from: number, to:number };
type Stack = string[];
type Stacks = Stack[];

const originalStacks: Stacks = [
    [ "N", "D", "M", "Q", "B", "P", "Z"],
    [ "C", "L", "Z", "Q", "M", "D", "H", "V"],
    [ "Q", "H", "R", "D", "V", "F", "Z", "G" ],
    [ "H", "G", "D", "F", "N"],
    [ "N", "F", "Q" ],
    [ "D", "Q", "V", "Z", "F", "B", "T" ],
    [ "Q", "M", "T", "Z", "D", "V", "S", "H"], 
    [ "M", "G", "F", "P", "N", "Q"],
    [ "B", "W", "R", "M"]
];

const moves: Move[] = lines.filter(line => line.startsWith("move")).map(line => line.replaceAll(/move |from |to /g, "")).
    map(numbers => numbers.split(" ")).map(n => ({ times: Number(n[0]), from: Number(n[1]) - 1, to: Number(n[2]) - 1 }));

function moveOne1(stacks: Stacks, { times, from, to }: Move) : void { 
    moveCrates(stacks[from], stacks[to], times);
};

function moveOne2(stacks: Stacks, { times, from, to }: Move) { 
    const temp: Stack = [];
    moveCrates(stacks[from], temp, times);
    moveCrates(temp, stacks[to], times);
};

function moveCrates(stackFrom: Stack, stackTo: Stack, times: number) : void {
    for (let i = 0; i < times; i++) {
        stackTo.push(stackFrom.pop()!) 
    }
}

function run(moveOne: (stacks: Stacks, move: Move) => void) {
    const stacks = originalStacks.map(elm => [...elm]); // deep copy
    moves.forEach(move => moveOne(stacks, move)); 
    console.debug(stacks.map(stack => stack.pop()!).join(""));
};

run(moveOne1); // QGTHFZBHV
run(moveOne2); // MGDMPSZTM
