import { readFileSync } from 'fs';

const inputContent = readFileSync("inputs/14_input.txt", 'utf-8');
const sampleContent = readFileSync("inputs/14_sample.txt", 'utf-8');

function run(content: string, hasFloor: boolean) {
    const lines = content.split('\n');
    lines.pop(); // Remove last empty line

    const input = lines.map(line => line.split(" ->").map(coord => ([...coord.split(",")].map(Number))).map(([x, y]) => [y, x]));
    const all = input.flat();
    let height = Math.max(...all.map(([y, _x]) => y)) + 1;
    let width = Math.max(...all.map(([_y, x]) => x)) + 1;
    let xoffset = 0;
    if (hasFloor) { 
        height += 2; 
        xoffset = height;
        width += xoffset * 2; 
    }
    const free = new Array(height) as boolean[][];
    for (let i = 0; i < height; i++) free[i] = new Array(width).fill(true);
    if (hasFloor) free[height - 1] = new Array(width).fill(false);

    for (const path of input as any[]) {
        for (let i = 1; i < path.length; i++) {
            const [yini, xini] = path[i-1];
            const [yend, xend] = path[i];
            const yinc = Math.sign(yend - yini), xinc = Math.sign(xend - xini);
    
            for (let y = yini, x = xini; y !== yend || x !== xend; y += yinc, x += xinc) {
                free[y][x - xoffset] = false;
            }
            free[yend][xend - xoffset] = false;
        }
    }

    outer: for (var total = 0; free[0][500 - xoffset] ; total++) {
        let y = 0, x = 500 - xoffset;
        while (true) {
            if (y == height-1) break outer;
            if (free[y+1][x]) {
                y++;
            } else if (free[y+1][x-1]) {
                y++; x--;
            } else if (free[y+1][x+1]) {
                y++; x++;
            } else break;
        }
        free[y][x] = false;
    }
    console.debug(total);    
};

run(sampleContent, false); // 24 (sample)
run(sampleContent, true);  // 93 (sample)
run(inputContent, false);  // 799
run(inputContent, true);   // 29076
