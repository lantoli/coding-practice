import { readFileSync } from 'fs';

const inputContent = readFileSync("inputs/18_input.txt", 'utf-8');
const sampleContent = readFileSync("inputs/18_sample.txt", 'utf-8');

const incs = [ [1, 0, 0], [-1, 0, 0], [0, 1, 0], [0, -1, 0], [0, 0, 1], [0, 0, -1] ]

function fromCube(max: number, cube: number) : number[] {
    const x = Math.floor(cube / (max * max))
    const y = Math.floor((cube - x * max * max) / max)
    const z = cube - x * max * max - y * max
    return [x, y, z]
}

function toCube(max: number, [x, y, z]: [number, number, number]) : number {
    return x * max * max + y * max + z
}

function getFree(cubes: Set<number>, max: number) : Set<number> {
    const free = new Set<number>()
    const states: number[] = []
    for (let y = 0; y < max; y++) for (let z = 0; z < max; z++) for (const x of [0, max-1]) 
        if (!cubes.has(toCube(max, [x, y, z]))) states.push(toCube(max, [x, y, z]))
    for (let x = 0; x < max; x++) for (let z = 0; z < max; z++) for (const y of [0, max-1]) 
        if (!cubes.has(toCube(max, [x, y, z]))) states.push(toCube(max, [x, y, z]))
    for (let x = 0; x < max; x++) for (let y = 0; y < max; y++) for (const z of [0, max-1]) 
        if (!cubes.has(toCube(max, [x, y, z]))) states.push(toCube(max, [x, y, z]))

    while (states.length > 0) {
        const cube = states.pop()!
        if (free.has(cube)) continue
        free.add(cube)
        const [x, y, z] = fromCube(max, cube)
        for (const [xinc, yinc, zinc] of incs) {
            const xx = x + xinc, yy = y + yinc, zz = z + zinc
            if (xx < 1 || xx > max-1 || yy < 1 || yy > max-1 || zz < 1 || zz > max-1) continue
            const cube = toCube(max, [xx, yy, zz])
            if (!free.has(cube) && !cubes.has(cube)) states.push(cube)
        }
    }
    return free
}

function run(content: string) {
    const lines = content.split('\n');
    lines.pop(); // Remove last empty line

    const input = lines.map(line => line.split(",").map(Number))
    const max = input.flat().reduce((a, b) => Math.max(a, b)) + 3
    const cubes = new Set(input.map(([x, y, z]) => toCube(max, [x+1, y+1, z+1])))
    const free = getFree(cubes, max)

    let total1 = 0, total2 = 0
    for (let cube of cubes) {
        const [x, y, z] = fromCube(max, cube)
        total1 += 6 - incs.filter(([xinc, yinc, zinc]) => cubes.has(toCube(max, [x + xinc, y + yinc, z + zinc]))).length
        total2 += incs.filter(([xinc, yinc, zinc]) => free.has(toCube(max, [x + xinc, y + yinc, z + zinc]))).length
    }

    console.debug(total1, total2)
};

run(sampleContent); // 64 58 (sample)
run(inputContent); // 4242 2428

