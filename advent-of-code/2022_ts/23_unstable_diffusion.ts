import { readFileSync } from 'fs'
import assert from 'assert'

const inputContent = readFileSync("inputs/23_input.txt", 'utf-8')
const sampleContent = readFileSync("inputs/23_sample.txt", 'utf-8')

const dirs = [
    [[-1, 0], [-1, -1], [-1, 1]],
    [[1, 0] , [1, -1] , [1, 1]],
    [[0, -1], [-1, -1], [1, -1]],
    [[0, 1] , [-1, 1] , [1, 1]],
]
const dirsFlat = dirs.flat()
const roundsPart1 = 10
const maxRounds = 2000
const inc = maxRounds * 2 // safety margin to avoid negative coordinates

function run(content: string) {
    const lines = content.split('\n')
    lines.pop() // Remove last empty line
    
    assert (lines.length === lines[0].length)
    const dim = lines.length
    const mod = dim + inc * 3
    const fromCoord = (y: number, x: number) => (y + inc) * mod + x + inc
    const toCoord = (n: number) => [ Math.floor(n / mod) - inc, (n % mod) - inc ]
    
    let prevMap = new Set<number>()
    const map = new Set<number>()
    for (let y = 0; y < dim; y++) for (let x = 0; x < dim; x++) {
        if (lines[y][x] === '#') map.add(fromCoord(y, x))
    }

    for (let round = 1, dirIndex = 0; round <= maxRounds; round++, dirIndex = (dirIndex + 1) % dirs.length) {
        const proposed = new Map<number, number>()
        for (const current of map) {
            const [y, x] = toCoord(current)
            if (dirsFlat.some(([yinc, xinc]) => map.has(fromCoord(y + yinc, x + xinc)))) {
                for (let i = 0; i < dirs.length; i++) {
                    const dir = dirs[(dirIndex + i) % dirs.length]
                    if (dir.every(([yinc, xinc]) => !map.has(fromCoord(y + yinc, x + xinc)))) {
                        const next = fromCoord(y + dir[0][0], x + dir[0][1])
                        proposed.set(next, proposed.has(next) ? -1 : current)
                        break
                    }
                }
            } 
        }
        for (const [next, current] of proposed) if (current !== -1) {
            map.delete(current)
            map.add(next)
        }

        if (round === roundsPart1) {
            let ymin = Infinity, ymax = 0, xmin = Infinity, xmax = 0
            for (const elf of map) {
                const [y, x] = toCoord(elf)
                ymin = Math.min(ymin, y)
                ymax = Math.max(ymax, y)
                xmin = Math.min(xmin, x)
                xmax = Math.max(xmax, x)
            }
            let total = 0
            for (let y = ymin; y <= ymax; y++) for (let x = xmin; x <= xmax; x++) {
                if (!map.has(fromCoord(y, x))) total++;
            }
            console.debug(total)
        }

        if ([...map].every(elf => prevMap.has(elf))) {
            console.debug(round)
            break
        } else prevMap = new Set(map)
    }
}

run(sampleContent) // 110, 20 (sample)
run(inputContent) // 4000, 1040
