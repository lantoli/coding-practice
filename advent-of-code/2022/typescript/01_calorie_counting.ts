import { readFileSync } from 'fs';

const inputContent = readFileSync("inputs/01_input.txt", 'utf-8');
const lines = inputContent.split('\n');

function run(count: number) {
    let total = 0;
    let list: number[] = []; // We could use a Heap but problem is small

    for (let line of lines) {
        const val = Number(line);
        total += val;
        if (val == 0) {
            list.push(total);
            total = 0;
        }
    }
    const top = list.sort((a, b) => b - a).slice(0, count);
    console.debug(top.reduce((a, b) => a + b)); 
};

run(1); // 67027
run(3); // 197291
