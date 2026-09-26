import { readFileSync } from 'fs';

const inputContent = readFileSync("inputs/04_input.txt", 'utf-8');
const lines = inputContent.split('\n');
lines.pop(); // Remove last empty line

type Coord = { left: number, right: number };
type CoordPair = [Coord, Coord];

const hasOverlap1 = ([r1, r2]: CoordPair) => r1.right >= r2.right;
const hasOverlap2 = ([r1, r2]: CoordPair) => r1.right >= r2.left;

function leftBiggerFirst([r1, r2]: CoordPair) : CoordPair {
    return (r1.left < r2.left) || (r1.left == r2.left && r1.right >= r2.right) ? [r1, r2] : [r2, r1];
}

function coords(line:string) : CoordPair {
    return line.split(',').map(elm => elm.split('-')).map(elm => ({left: Number(elm[0]), right: Number(elm[1])})) as CoordPair;
}

function run(hasOverlap : (coords: CoordPair) => boolean) {
   const result = lines.map(coords).map(leftBiggerFirst).filter(hasOverlap).length; 
   console.debug(result);
};

run(hasOverlap1); // 657
run(hasOverlap2); // 938
