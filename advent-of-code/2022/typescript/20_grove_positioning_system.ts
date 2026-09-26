import { readFileSync } from 'fs'

const inputContent = readFileSync("inputs/20_input.txt", 'utf-8')
const sampleContent = readFileSync("inputs/20_sample.txt", 'utf-8')

function run(content: string, decryptionKey: number, rounds: number) {
    const lines = content.split('\n')
    lines.pop() // Remove last empty line

    let list = lines.map((line, index) => [Number(line) * decryptionKey, index])
    const mod = list.length
    
    for (let round = 0; round < rounds; round++) {
        for (let i = 0; i < list.length; i++) {
            const ini = list.findIndex(([_, index]) => index === i)
            const n = list[ini][0] % (mod - 1)
            const end = (ini + n + mod - 1) % (mod - 1)
            const dir = Math.sign(end - ini)
            for (let pos = ini; pos != end; pos += dir) {
                [list[pos], list[pos + dir]] = [list[pos + dir], list[pos]]
            }
        }
    }

    const idx0 = list.findIndex(([n, _]) => n === 0)
    const sum = [1000, 2000, 3000].map(n => list[(n + idx0) % mod][0]).reduce((a, b) => a + b)
    console.debug(sum)
}

run(sampleContent, 1, 1);               // 3 (sample)
run(sampleContent, 811589153, 10);      // 1623178306 (sample)
run(inputContent, 1, 1);                // 3700
run(inputContent, 811589153, 10);       // 10626948369382
