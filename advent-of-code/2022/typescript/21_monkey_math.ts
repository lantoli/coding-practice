// NOTE: parte 2 is very slow so I solved manually the equation, it won't work for other inputs

import { readFileSync } from 'fs'

const inputContent = readFileSync("inputs/21_input.txt", 'utf-8')
const sampleContent = readFileSync("inputs/21_sample.txt", 'utf-8')

type Monkey = { name: string, expr: string, value?: number }

function getMonkeys(content: string) : Map<string, Monkey> {
    const lines = content.split('\n')
    lines.pop() // Remove last empty line

    const map = new Map<string, Monkey>()
    for (const line of lines) {
        const [name, expr] = line.split(': ')
        const value = parseInt(expr)
        const monkey = { name, expr, value: isNaN(value) ? undefined : value}
        map.set(name, monkey,)
    };
    return map
}

function getValue(map: Map<string, Monkey>, monkey: Monkey) : number {
    if (monkey.value !== undefined) return monkey.value
    const [exp1, op, exp2] = monkey.expr.split(" ")
    let ret = 0
    switch(op) {
        case "+": ret = getValue(map, map.get(exp1)!) + getValue(map, map.get(exp2)!); break
        case "-": ret =  getValue(map, map.get(exp1)!) - getValue(map, map.get(exp2)!); break
        case "*": ret = getValue(map, map.get(exp1)!) * getValue(map, map.get(exp2)!); break
        case "/": ret = Math.floor(getValue(map, map.get(exp1)!) / getValue(map, map.get(exp2)!)); break
        default: throw new Error("Unknown operator: " + op)
    }
    return monkey.value = ret
}

function run1(content: string) {
    const map = getMonkeys(content)
    const root = map.get("root")!
    console.debug(getValue(map, root))
}

function run2(content: string) {
    const map = getMonkeys(content)
    const root = map.get("root")!
    const [exp1, _op, exp2] = root.expr.split(" ")
    const val2 = getValue(map, map.get(exp2)!)
    const snapshot = new Map([...map.entries()].map(([name, monkey]) => [name, monkey.value]))

    for (let num = 0; ; num ++) {
        map.get("humn")!.value = num
        if (getValue(map, map.get(exp1)!) === val2) {
            console.debug(num)
            return
        }
        for (const monkey of map.values()) monkey.value = snapshot.get(monkey.name)
    }
}

function run2b() {
    for (let num = 3887609741064; ; num ++) {
        const val = fnPart2(num)
        const cmp = 32310041242752
        if (val === cmp) {
            console.debug(num)
            return
        }
    }
}

function drawFunction() {
    // ((27)) * (((6007094388653)) - (Math.floor(((((732)) + (((787)) + (((Math.floor((((((577)) + (Math.floor(((Math.floor(((((2)) * ((Math.floor(((((2)) * ((((2)) * ((Math.floor((((((76)) + (((Math.floor((((534)) + (Math.floor(((((2)) * ((((Math.floor(((((Math.floor((((996)) + (Math.floor((((623)) + ((((476)) + (((((Math.floor(((((5)) * ((Math.floor(((((28)) * ((Math.floor(((((622)) + (Math.floor(((((5)) * ((Math.floor((((533)) + ((((humn)) - ((876))) * ((7)))) / ((2)))) - ((475)))) + ((690))) / ((5))))) + ((591))) / ((11)))) - ((150)))) - ((670))) / ((2)))) + ((471)))) + ((432))) / ((12)))) - ((26))) * ((2))) - ((393))) * ((9)))) + ((570)))) / ((2))))) / ((2)))) - ((148))) * ((3))) - ((128))) / ((2)))) + ((31))) * ((2))) - ((519)))) + ((834))) / ((3))))) / ((2)))) - ((448))) * ((3)))) + ((94))) + ((397))) / ((4)))) - ((572)))) + ((805)))) - ((267))) / ((9)))) - ((451)))) + ((414))) / ((2)))) - ((157))) / ((2))))) * ((12))) - ((723))) / ((9)))) + ((410))) * ((4))))) + ((323))) / ((6)))))
    // Simplifying with some decimal errors: humn = 3887609741064,49
    const map = getMonkeys(inputContent)
    const root = map.get("root")!
    const [exp1, _op, exp2] = root.expr.split(" ")
    
    map.get("humn")!.value = NaN

    const val1 = getValue(map, map.get(exp1)!)
    const val2 = getValue(map, map.get(exp2)!)

    console.debug(draw(map, exp1))
}

function draw(map: Map<string, Monkey>, name: string) : string {
    if (name.includes("humn")) return `(${name})`
    const monkey = map.get(name)!    
    if (monkey.value !== undefined && !Number.isNaN(monkey.value)) return `(${monkey.value})`
    const [exp1, op, exp2] = monkey.expr.split(" ")
    const draw1 = map.has(exp1) ? draw(map, exp1) : exp1
    const draw2 = map.has(exp2) ? draw(map, exp2) : exp2
    
    let ret = `(${draw1}) ${op} (${draw2})`
    if (op === "/") ret = `Math.floor(${ret})`
    return ret
}

// Got from drawFunction
function fnPart2(humn: number) {
    return ((27)) * (((6007094388653)) - (Math.floor(((((732)) + (((787)) + (((Math.floor((((((577)) + (Math.floor(((Math.floor(((((2)) * ((Math.floor(((((2)) * ((((2)) * ((Math.floor((((((76)) + (((Math.floor((((534)) + (Math.floor(((((2)) * ((((Math.floor(((((Math.floor((((996)) + (Math.floor((((623)) + ((((476)) + (((((Math.floor(((((5)) * ((Math.floor(((((28)) * ((Math.floor(((((622)) + (Math.floor(((((5)) * ((Math.floor((((533)) + ((((humn)) - ((876))) * ((7)))) / ((2)))) - ((475)))) + ((690))) / ((5))))) + ((591))) / ((11)))) - ((150)))) - ((670))) / ((2)))) + ((471)))) + ((432))) / ((12)))) - ((26))) * ((2))) - ((393))) * ((9)))) + ((570)))) / ((2))))) / ((2)))) - ((148))) * ((3))) - ((128))) / ((2)))) + ((31))) * ((2))) - ((519)))) + ((834))) / ((3))))) / ((2)))) - ((448))) * ((3)))) + ((94))) + ((397))) / ((4)))) - ((572)))) + ((805)))) - ((267))) / ((9)))) - ((451)))) + ((414))) / ((2)))) - ((157))) / ((2))))) * ((12))) - ((723))) / ((9)))) + ((410))) * ((4))))) + ((323))) / ((6)))))
}

run1(sampleContent); // 152 (sample)
run2(sampleContent); // 301 (sample)
run1(inputContent); // 194501589693264
run2b() // 3887609741189
