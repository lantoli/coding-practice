import { readFileSync, stat } from 'fs';

const inputContent = readFileSync("inputs/08_input.txt", 'utf-8');
const lines = inputContent.split('\n');
lines.pop(); // Remove last empty line

const grid = lines.map(line => [...line].map(ch => Number(ch)))
const height = grid.length, width = grid[0].length;
const incs = [ [1, 0], [0, 1], [-1, 0], [0, -1] ];

function isVisible(y: number, x: number) : boolean {
    return incs.some(([yinc, xinc]) => canReachSide(y, x, yinc, xinc));
}

function canReachSide(y: number, x: number, yinc: number, xinc: number) : boolean {
    return untilBorder(y, x, yinc, xinc).every(([yy, xx]) => grid[yy][xx] < grid[y][x]);    
}

function allTreesMultiply(y: number, x: number) : number {
    return incs.map(([yinc, xinc]) => treesReach(y, x, yinc, xinc)).reduce((a, b) => a * b);
}

function treesReach(y: number, x: number, yinc: number, xinc: number) : number 
{
    let score = 0;
    for (let [yy, xx] of untilBorder(y, x, yinc, xinc)) {
        score++;
        if (grid[yy][xx] >= grid[y][x]) break;
    };
    return score;
}

function untilBorder(y: number, x: number, yinc: number, xinc: number) : number[][] {
    const isBorder = (y: number, x: number) => y == 0 || y == height - 1 || x == 0 || x == width - 1;
    const ret : number[][] = [];
    while (!isBorder(y, x)) {
        y += yinc; x += xinc;
        ret.push([y, x]);
    };
    return ret;
}

function run() {
    let part1 = 0, part2 = 0;
    for (let y = 0; y < height; y++) // not efficient at all but problem size is tiny
        for (let x = 0; x < width; x++) {
            if (isVisible(y, x)) part1++;
            part2 = Math.max(part2, allTreesMultiply(y, x));
        }
    console.debug(part1, part2);
};

run(); // 1801 209880
