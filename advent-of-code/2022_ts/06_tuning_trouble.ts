import { readFileSync } from 'fs';

const inputContent = readFileSync("inputs/06_input.txt", 'utf-8');

function position(data: String, len: number) : number {
    for (let i = len; i <= data.length; i++) {
        if (new Set(data.substring(i-len, i)).size == len) return i;
    }
    throw new Error("not found");
}

function run(len: number) {
    console.debug(position(inputContent, len)); 
};

run(4); // 1766
run(14); // 2383
