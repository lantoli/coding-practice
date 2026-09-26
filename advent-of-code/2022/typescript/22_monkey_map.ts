// Note: wrapInput assumes the face layout in my input, it may fail for other input layouts. 
// I created the die in paper to figure out the wrapping rules.
import { readFileSync } from 'fs'
import assert from 'assert';

const inputContent = readFileSync("inputs/22_input.txt", 'utf-8')
const sampleContent = readFileSync("inputs/22_sample.txt", 'utf-8')

type Move = { dir: number, ycube: number, xcube: number, ypos: number, xpos:number }

const right = 0, down = 1, left = 2, up = 3
const incs = [ [0, 1], [1, 0], [0, -1], [-1, 0]  ];

const pwd = (y: number, x: number, dir: number) => 1000 * (y+1) + 4 * (x+1) + dir

function getBoard(content: string, size: number) {
  const lines = content.split('\n')
  lines.pop() // Remove last empty line
  const grid: string[] = lines.slice(0, lines.length - 2)
  const path = lines[lines.length - 1]

  assert(grid.length % size === 0)
  const faceRows = Math.floor(grid.length / size)
  const faceCols = Math.floor(12 / faceRows)
  const rows = faceRows * size
  const cols = faceCols * size
  for (let i = 0; i < grid.length; i++) {
    assert(grid[i].length % size === 0)
    assert(grid[i].length <= cols)
    grid[i] += ' '.repeat(cols - lines[i].length)
    assert(grid[i].length == cols)
  }
  return { grid, faceRows, faceCols, rows, cols, path }
}

function wrapSample(m: Move, size: number) : Move {
  const { dir, ycube, xcube, ypos, xpos } = m
  if (ycube == 1 && xcube == 3 && xpos == 0 && dir === 0)
    return { dir: 1, ycube: ycube + 1, xcube, ypos: xpos, xpos: size - ypos - 1 }
  
  if (ycube == 0 && xcube == 2 && ypos == 0 && dir == 1) 
    return { dir: 3, ycube: ycube + 1, xcube: 0, ypos: size - 1, xpos: size - xpos - 1 }

  if (ycube == 0 && xcube == 1 && ypos == size - 1 && dir == 3) 
    return { dir: 0, ycube, xcube: xcube + 1, ypos: xpos, xpos: 0 }

  throw new Error("Wrapping not supported")
}

function wrapInput(m: Move, size: number) : Move {
  const { dir, ycube, xcube, ypos, xpos } = m
  const reverse = (x: number) => size - x - 1
  const opposite = (x: number) => x === 0 ? border : x === border ? 0 : Number.NaN
  const border = size - 1
  const dirL = { dir: (dir + 3) % 4, xcube, ypos: reverse(xpos), xpos: ypos }
  const dirR = { dir: (dir + 1) % 4, ycube, ypos: xpos, xpos: reverse(ypos) }
  const dir180 = { dir: (dir + 2) % 4, ypos: reverse(ypos), xpos: reverse(xpos) }

  switch(pwd(ycube, xcube, dir)) {
    case pwd(3, 2, up   ): return { ...m, xcube: 0 }
    case pwd(0, 0, down ): return { ...m, xcube: 2 }

    case pwd(3, 1, right): return { ...dirL, ycube: 2 }
    case pwd(1, 0, left ): return { ...dirL, ycube: 2 } 
    case pwd(1, 2, right): return { ...dirL, ycube: 0 }
    case pwd(3, 2, left ): return { ...dirL, ycube: 0 } 

    case pwd(3, 1, up   ): return { ...dirR, xcube: 0, xpos: opposite(ypos) }
    case pwd(3, 1, down ): return { ...dirR, xcube: 0 }
    case pwd(1, 0, up   ): return { ...dirR, xcube: 1 }
    case pwd(1, 2, down ): return { ...dirR, xcube: 1 }

    case pwd(2, 2, left ): return { ...dir180, ycube: 0, xcube: 1, xpos: opposite(xpos) } 
    case pwd(2, 2, right): return { ...dir180, ycube: 0, xcube: 2 } 
    case pwd(0, 0, left ): return { ...dir180, ycube: 2, xcube: 0 } 
    case pwd(0, 0, right): return { ...dir180, ycube: 2, xcube: 1 } 
    
    default: throw new Error("Wrapping not supported")
  }
}

function run(content: string, size: number, wrap?: (m: Move, size: number) => Move) {
  const { grid, faceRows, faceCols, rows, cols, path } = getBoard(content, size)
  let ylast = 0, xlast = grid[0].indexOf('.'), dirlast = 0 // DEJAR ESTE

  for (let i = 0; i < path.length; ) switch(path[i]) {
    case 'L': dirlast = (dirlast + 3) % 4; i++; break;
    case 'R': dirlast = (dirlast + 1) % 4; i++; break;
    default: 
      for (var moves = 0; i < path.length && path[i] !== 'L' && path[i] !== 'R'; i++) {
        moves = moves * 10 + path[i].charCodeAt(0) - '0'.charCodeAt(0)
      }
      for (let move = 0, y = ylast, x = xlast, dir = dirlast; move < moves; move++, ylast = y, xlast = x, dirlast = dir) {
        do {
          y += incs[dir][0]; x += incs[dir][1]
          let overflow = false
          if (y == -1) { y = rows - 1; overflow = true }
          if (y == rows) { y = 0; overflow = true }
          if (x == -1) { x = cols - 1; overflow = true }
          if (x == cols) { x = 0; overflow = true }

          if (typeof wrap !== "undefined") {
            const ycube = Math.floor(y / size), xcube = Math.floor(x / size)
            const ypos = y % size, xpos = x % size

            if (overflow || grid[y][x] === ' ')  {
              const info = wrap({ dir, ycube, xcube, ypos, xpos }, size)
              y = info.ycube * size + info.ypos
              x = info.xcube * size + info.xpos
              dir = info.dir
            }
            assert (grid[y][x] !== ' ')
          }
        } while (grid[y][x] === ' ')
        if (grid[y][x] === '#') break
      }
      break;
  }
  console.debug(pwd(ylast, xlast, dirlast))
}

run(sampleContent, 4);                // 6032 (sample)
run(sampleContent, 4, wrapSample);    // 5031 (sample)
run(inputContent, 50);                // 109094
run(inputContent, 50, wrapInput);     // 53324
