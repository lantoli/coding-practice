// Part2 takes 10m
// Use: export NODE_OPTIONS="--max-old-space-size=16384"
import { readFileSync } from 'fs'

const inputContent = readFileSync("inputs/19_input.txt", 'utf-8')
const sampleContent = readFileSync("inputs/19_sample.txt", 'utf-8')

enum res { ore, clay, obsidian, geode }
const len = Object.keys(res).length / 2
type State = { resourcesNumber: number, robotsNumber: number, minute: number }

const mod = 1000
const mod4 = mod * mod * mod * mod

function fromTuple([a, b, c, d]: number[]): number {
  return (a * mod * mod * mod + b * mod * mod + c * mod + d)
}

function toTuple(n: number): number[] {
  const a = Math.floor(n / (mod * mod * mod))
  const b = Math.floor(n / (mod * mod)) % mod
  const c = Math.floor(n / mod) % mod
  const d = n % mod
  return [a, b, c, d]
}

function robotCostFromBlueprint(line: string): number[][] {
  const oreRobot = line.match(/Each ore robot costs (\d+) ore./)!
  const clayRobot = line.match(/Each clay robot costs (\d+) ore./)!
  const obsidianRobot = line.match(/Each obsidian robot costs (\d+) ore and (\d+) clay./)!
  const geodeRobot = line.match(/Each geode robot costs (\d+) ore and (\d+) obsidian./)!

  const robots : number[][] = new Array(len).fill(0).map(() => new Array(len).fill(0));
  robots[res.ore][res.ore] = parseInt(oreRobot[1])
  robots[res.clay][res.ore] = parseInt(clayRobot[1])
  robots[res.obsidian][res.ore] = parseInt(obsidianRobot[1])
  robots[res.obsidian][res.clay] = parseInt(obsidianRobot[2])
  robots[res.geode][res.ore] = parseInt(geodeRobot[1])
  robots[res.geode][res.obsidian] = parseInt(geodeRobot[2])
  return robots
}

function getGeodes(robotCost: number[][], minutes: number): number {
  const states: State[] = [ {resourcesNumber: 0, robotsNumber: fromTuple([1, 0, 0, 0]), minute: 0 } ]
  const visited = new Map<number, State>()
  let maxGeodes = 0

  while (states.length > 0) {
    const state = states.pop()!
    let { resourcesNumber, robotsNumber, minute } = state
    let resources = toTuple(resourcesNumber)
    let robots = toTuple(robotsNumber)
    
    if (visited.size > 10_000_000) visited.clear() // to avoid out of memory errors
    let key = robotsNumber * mod4 + resourcesNumber
    const v = visited.get(key)
    if (typeof v !== "undefined" && state.minute >= v.minute) continue
    visited.set(key, state)

    for (let i = 0; i < len; i++) { // each next state is a robot being built (if possible)
      let rounds = 0
      for (let j = 0; j < len; j++)  {
        const cost = robotCost[i][j];
        if (resources[j] < cost) {
          if (robots[j] > 0) {
            rounds = Math.max(rounds, Math.ceil((cost - resources[j]) / robots[j]))
          } else {
            rounds = Infinity
          }
        }
      }
      if (Number.isFinite(rounds)) {
        rounds++
        const newMinute = minute + rounds
        if (newMinute < minutes) {
          const newRobots = [...robots]
          newRobots[i]++
          const newResources = [...resources]
          for (let j = 0; j < len; j++) {
            newResources[j] += rounds * robots[j] - robotCost[i][j]
          }
          const newState = { resourcesNumber: fromTuple(newResources), robotsNumber: fromTuple(newRobots), minute: newMinute }
          states.push(newState)
          maxGeodes = Math.max(maxGeodes, newRobots[res.geode])
        } else {
          maxGeodes = Math.max(maxGeodes, resources[res.geode] + (minutes - minute) * robots[res.geode])
        } 
      }
    }
  }
  return maxGeodes
}

function run1(content: string) {
  const lines = content.split('\n')
  lines.pop() // Remove last empty line
  const minutes = 24

  const geodes = lines.map(robotCostFromBlueprint).map(cost => getGeodes(cost, minutes))
  const ret = geodes.map((geode, idx) => (idx + 1) * geode).reduce((a, b) => a + b)
  console.debug(ret)
}

function run2(content: string) {
  const lines = content.split('\n')
  lines.pop() // Remove last empty line
  const minutes = 32

  const geodes = lines.map(robotCostFromBlueprint).slice(0, 3).map(cost => getGeodes(cost, minutes))
  const ret = geodes.reduce((a, b) => a * b)
  console.debug(ret)
}

run1(sampleContent); // 33 (sample)
run2(sampleContent); // 3348 (sample)
run1(inputContent); // 1616
run2(inputContent); // 8990
