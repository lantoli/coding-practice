import { readFileSync, stat } from 'fs';

const inputContent = readFileSync("inputs/09_input.txt", 'utf-8');
const lines = inputContent.split('\n');
lines.pop(); // Remove last empty line

const input: [string, number][] = lines.map(line => line.split(" ")).map(([dir, num]) => [dir, Number(num)]);

type Coord = {y: number, x: number};

function moveTail(tail: Coord, head: Coord) : Coord {
    const ydist = Math.abs(head.y - tail.y), xdist = Math.abs(head.x - tail.x);
    const yinc = Math.sign(head.y - tail.y), xinc = Math.sign(head.x - tail.x);
    if  (ydist == 2 || xdist == 2) {
        tail = { y: tail.y + yinc, x: tail.x + xinc };
    }
    return tail;
}

function run(knotCount: number) {
    const moveHead = (yinc: number, xinc: number) => knots[0] = { y: knots[0].y + yinc, x: knots[0].x + xinc };
    const knots: Coord[] = [];
    for (let i = 0; i < knotCount; i++) knots.push({ y: 0, x: 0});
    let tailHist = [ knots[knotCount - 1] ];
    
    for (let [dir, num] of input) for (let i = 0; i < num; i++) {
        switch(dir) {
            case "R": moveHead(0, 1); break;
            case "L": moveHead(0, -1); break;
            case "D": moveHead(1, 0); break;
            case "U": moveHead(-1, 0); break;
            default: throw new Error("unknown");
        }
        for (let i = 1; i < knotCount; i++) knots[i] = moveTail(knots[i], knots[i-1]);

        const tail = knots[knotCount-1];
        if (!tailHist.some(elm => elm.x === tail.x && elm.y === tail.y)) {
            tailHist.push(tail);
        }
    }
    console.log(tailHist.length);
};

run(2); // 6030 
run(10); // 2545
