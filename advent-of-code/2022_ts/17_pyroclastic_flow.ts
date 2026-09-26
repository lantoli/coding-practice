import { readFileSync } from 'fs'

const inputContent = readFileSync("inputs/17_input.txt", 'utf-8')
const sampleContent = readFileSync("inputs/17_sample.txt", 'utf-8')

const shapes = [
    [ [true, true, true, true] ],
    [ [false, true, false], [true, true, true], [false, true, false] ],
    [ [false, false, true], [false, false, true], [true, true, true] ], 
    [ [true],[true], [true], [true] ],
    [ [true, true], [true, true] ]
]

const safetyRows = 30 // enough rows to get the same result

function allowed(board: number[], shape: boolean[][], ypos: number, xpos: number) : boolean {
    for (let yshape = 0; yshape < shape.length; yshape++) for (let xshape = 0; xshape < shape[0].length; xshape++) {
        if (shape[yshape][xshape] && (board[ypos - yshape] & 1 << (xpos+xshape))) return false
    }
    return true;
}

function putPiece(board: number[], shape: boolean[][], ypos: number, xpos: number) : void {
    for (let yshape = 0; yshape < shape.length; yshape++) for (let xshape = 0; xshape < shape[0].length; xshape++) {
        if (shape[yshape][xshape]) board[ypos - yshape] |= 1 << xpos + xshape
    }
}

const highestRock = (board: number[]): number => {
    for (let y = board.length - 1; y >= 0; y--) if (board[y] > 0) return y
    throw new Error("shouldnt happen")
}

function run(content: string, rocks: number) {
    const lines = content.split('\n')
    lines.pop() // Remove last empty line
    const moves = [...lines[0]]

    const width = 7
    let board: number[] = [255] // wall (7 bits set)
    let moveIndex = 0
            
    const mod = Math.max(shapes.length, moves.length) + 1
    const cycles = new Map<number, { safetyBoard: string, count: number, highest: number }>();
    let extra = 0

    for (let count = 0, shapeIndex = 0; count < rocks; count++, shapeIndex = (shapeIndex + 1) % shapes.length) {
        const shape  = shapes[shapeIndex]
        const highest = highestRock(board)
        const safetyBoard = JSON.stringify(board.slice(highest + 1 - safetyRows, highest + 1))
        let ypos = highest + 3 + shape.length
        let xpos = 2
        while (board.length <= ypos) board.push(0)

        const id = shapeIndex * mod + moveIndex
        const info = cycles.get(id)
        if (typeof info !== 'undefined' && info.safetyBoard === safetyBoard) {
            const diffCount = count - info.count
            const diffHighest = highest - info.highest
            const times = Math.floor((rocks - count) / diffCount)
            count += diffCount * times
            extra += diffHighest * times
        } else {
            cycles.set(id, { safetyBoard, count, highest: highest + extra })
        }
        
        while (true) { // one piece falling down
            const move = moves[moveIndex]
            moveIndex = (moveIndex + 1) % moves.length

            if (move === ">") {
                if (xpos + shape[0].length < width && allowed(board, shape, ypos, xpos+1)) xpos++;
            } else { // <
                if (xpos > 0 && allowed(board, shape, ypos, xpos-1)) xpos--;
            }
    
            if (allowed(board, shape, ypos - 1, xpos)) ypos--; else break
        }

        putPiece(board, shape, ypos, xpos)
    }

    return highestRock(board) + extra
}

console.debug(run(sampleContent, 2022))               // 3068 (sample)
console.debug(run(sampleContent, 1_000_000_000_000))    // 1514285714288 (sample)
console.debug(run(inputContent, 2022))                  // 3215
console.debug(run(inputContent, 1_000_000_000_000))     // 1575811209487
