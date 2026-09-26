import { readFileSync } from 'fs'

const inputContent = readFileSync("inputs/25_input.txt", 'utf-8')
const sampleContent = readFileSync("inputs/25_sample.txt", 'utf-8')

const mapFrom = new Map<string, number>([["0", 0], ["1", 1], ["2", 2], ["-", -1], ["=", -2]])

const mapTo = new Map<number, [string, number]>([
    [0, ["0", 0]], [1, ["1", 0]], [2, ["2", 0]],
    [3, ["=", 1]], [4, ["-", 1]]
])

const fromSnafu = (str: string) => [...str].reduce((acc, ch) => acc * 5 + mapFrom.get(ch)!, 0)

function toSnafu(n: number) : string {
    for (var result = ""; n > 0; ) {
        const [ch, carry] = mapTo.get(n % 5)!
        result = ch + result
        n = Math.floor(n / 5) + carry
    }
    return result
}

function run(content: string) {
    const lines = content.split('\n')
    lines.pop() // Remove last empty line
    console.debug(toSnafu(lines.map(fromSnafu).reduce((a, b) => a + b)))
}

run(sampleContent) // 2=-1=0 (sample)
run(inputContent) // 2-0-020-1==1021=--01
