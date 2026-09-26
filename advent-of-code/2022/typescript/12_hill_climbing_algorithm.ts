import { readFileSync } from 'fs';

const inputContent = readFileSync("inputs/12_input.txt", 'utf-8');
const sampleContent = readFileSync("inputs/12_sample.txt", 'utf-8');

type Coord = [number, number];
type State = { pos: Coord, count: number };

function run(content: string, tryAll: boolean) {
    const lines = content.split('\n');
    lines.pop(); // Remove last empty line

    let e: Coord = [0, 0];
    const states: State[] = [];
    const input = lines.map((line, y) => [...line].map((ch, x) => {
        if (ch === "S" || (ch === "a" && tryAll)) { 
            states.push({ pos: [y, x], count: 1 }); ch = "a"; 
        } else if (ch === "E") { 
            e = [y, x]; ch = "z"; 
        }
        return ch.charCodeAt(0) - "a".charCodeAt(0);
    }));

    const height = lines.length, width = lines[0].length;
    const visited = new Set<number>();
    while (states.length > 0) {
        const cur = states.shift()!;
        const [y, x] = cur.pos;
        const count = cur.count + 1;
        for (const [yinc, xinc] of [ [1, 0], [0, 1], [-1, 0], [0, -1] ]) {
            const yy = y + yinc, xx = x + xinc;
            if (yy >= 0 && yy < height && xx >= 0 && xx < width && (input[yy][xx] <= input[y][x] + 1)) {
                if (yy === e[0] && xx === e[1]) {
                    console.debug(cur.count);
                    return;
                }
                if (!visited.has(yy * width + xx)) {
                    visited.add(yy * width + xx);
                    states.push({ pos: [yy, xx], count });
                }
            } 
        }
    }
};

run(sampleContent, false); // 31 (sample)
run(sampleContent, true); // 29 (sample)
run(inputContent, false); // 440
run(inputContent, true); // 439
