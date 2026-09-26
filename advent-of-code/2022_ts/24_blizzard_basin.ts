import { readFileSync } from 'fs'

const inputContent = readFileSync("inputs/24_input.txt", 'utf-8')
const sample1Content = readFileSync("inputs/24_sample1.txt", 'utf-8')
const sample2Content = readFileSync("inputs/24_sample2.txt", 'utf-8')

const dirs = [ [-1, 0], [1, 0], [0, -1], [0, 1] ]
const dirChars = ["^", "v", "<", ">"]

class Info {
    rows: number; cols: number; rowsIn: number; colsIn: number; cycle: number; mod: number
    blizzards: number[][][]

    constructor(content: string) {
        const lines = content.split('\n')
        lines.pop() // Remove last empty line
    
        this.rows = lines.length; this.rowsIn = this.rows - 2
        this.cols = lines[1].length; this.colsIn = this.cols - 2
        this.cycle = this.rowsIn * this.colsIn
        this.mod = Math.max(this.cols, this.rows) + 2

        const first: number[][] = new Array(dirs.length).fill(0).map(() => [])
        for (let y = 0; y < this.rows; y++) for (let x = 0; x < this.cols; x++) {
            const dirIndex = dirChars.indexOf(lines[y][x])
            if (dirIndex != -1) first[dirIndex].push(this.fromCoord(y, x)) 
        }
        this.blizzards = [first]
        for (let idxHistory = 1; idxHistory < this.cycle; idxHistory++) {
            const prev = this.blizzards[idxHistory - 1]
            const current: number[][] = new Array(dirs.length).fill(0).map(() => [])
            this.blizzards.push(current)
            for (let idx in dirs) {
                const [yinc, xinc] = dirs[idx]
                for (const coord of prev[idx]) {
                    const [y, x] = this.toCoord(coord)
                    current[idx].push(this.fromCoord((y + yinc - 1 + this.rowsIn) % this.rowsIn + 1, (x + xinc - 1 + this.colsIn) % this.colsIn + 1))
                }
            }
        }
    }

    fromCoord = (y: number, x: number) => y * this.mod + x
    toCoord = (n: number) => [Math.floor(n / this.mod), n % this.mod]

    canCoord = (y: number, x: number) => y > 0 && y < this.rows - 1 && x > 0 && x < this.cols - 1 
    canBlizzards = (y: number, x: number, minute: number) => !this.blizzards[minute % this.cycle].some(dir => dir.includes(this.fromCoord(y, x)))
    
    path(from: number, to: number, minute: number) : number {
        const [yto, xto] = this.toCoord(to)
        for (let cur = new Set([from]), next = new Set<number>(); ; minute++, [cur, next] = [next, new Set()]) {
            for (const state of cur) {
                const [y, x] = this.toCoord(state)
                if (this.canBlizzards(y, x, minute)) next.add(this.fromCoord(y, x))
                for (const [yinc, xinc] of dirs) {
                    const [ynext, xnext] = [y + yinc, x + xinc]
                    if (yto === ynext && xto === xnext) return minute
                    if (this.canCoord(ynext,  xnext) && this.canBlizzards(ynext, xnext, minute)) next.add(this.fromCoord(ynext, xnext))
                } 
            }
        }
    }
}

function run(content: string) {
    const i = new Info(content)
    const ini = i.fromCoord(0, 1), end = i.fromCoord(i.rows - 1, i.cols - 2)
    const path1 = i.path(ini, end, 1)
    const path2 = i.path(end, ini, path1)
    const path3 = i.path(ini, end, path2)
    console.debug(path1, path3)
}

run(sample1Content) // 10, 29 (sample 1)
run(sample2Content) // 18, 54 (sample 2)
run(inputContent) // 255, 809
