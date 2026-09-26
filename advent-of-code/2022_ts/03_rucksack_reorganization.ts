import { readFileSync } from 'fs';

const inputContent = readFileSync("inputs/03_input.txt", 'utf-8');
const lines = inputContent.split('\n');
lines.pop(); // Remove last empty line

const lines1 = lines.map(line => [line.substring(0, line.length / 2), line.substring(line.length / 2)]);

const lines2: string[][] = [];
for (let i = 0; i < lines.length; i += 3) {
    lines2.push([lines[i], lines[i+1], lines[i+2]]);
}

function choose([r1, r2, r3]: string[]) : string {
    const [first, second, third] = [[...r1], [...r2], [...r3||r1]];
    return first.filter(letter => second.includes(letter) && third.includes(letter))[0];
}

function priority(letter: string) {
    const ascii = letter.charCodeAt(0);
    return ascii <= "Z".charCodeAt(0) ? 
        ascii - "A".charCodeAt(0) + 27 : 
        ascii - "a".charCodeAt(0) + 1; 
}

function run(data: string[][]) {
    const result = data.map(choose).map(priority).reduce((a, b) => a + b);
    console.debug(result);
};

run(lines1); // 7701
run(lines2); // 2644
