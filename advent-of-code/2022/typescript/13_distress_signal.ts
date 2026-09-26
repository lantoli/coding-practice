import { readFileSync } from 'fs';

const inputContent = readFileSync("inputs/13_input.txt", 'utf-8');
const sampleContent = readFileSync("inputs/13_sample.txt", 'utf-8');

function compare(left: any, right: any): number {
    if (typeof left === "string") left = JSON.parse(left);
    if (typeof right === "string") right = JSON.parse(right);
    if (Array.isArray(left) && Array.isArray(right)) {
        for (let i = 0; i < left.length && i < right.length; i++) {
            const ret = compare(left[i], right[i]);
            if (ret != 0) return ret;
        }
        return left.length - right.length;
    } else if (typeof left === "number" && typeof right === "number") {
        return left - right;
    } else if (Array.isArray(left)) {
        return compare(left, [right]);
    } else {
        return compare([left], right);
    }
}

function processTwo(content: string): [string, string][] {
    const lines = content.split('\n');
    lines.pop(); // Remove last empty line

    const ret: [string, string][] = [];
    for (let i = 0; i < lines.length; i += 3) {
        ret.push([lines[i], lines[i+1]]);
    }
    return ret;
}

function run1(content: string) {
    const ret: number[] = [0];
    processTwo(content).forEach(([left, right], idx) => {
        if (compare(left, right) <= 0) ret.push(idx + 1);
    });
    console.debug(ret.reduce((a, b) => a + b));
};

function run2(content: string) {
    const divs = ["[[2]]", "[[6]]"];
    const numbers = processTwo(content).flat();
    numbers.push(...divs);
    numbers.sort((a, b) => compare(a, b));
    const ret = divs.map(div => numbers.indexOf(div) + 1).reduce((a, b) => a * b);
    console.debug(ret);
};

run1(sampleContent); // 13 (sample)
run2(sampleContent); // 140 (sample)
run1(inputContent); // 6272
run2(inputContent); // 22288
