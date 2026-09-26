import { readFileSync, stat } from 'fs';

const inputContent = readFileSync("inputs/07_input.txt", 'utf-8');
const lines = inputContent.split('\n');
lines.pop(); // Remove last empty line

const state = {
    current: "/",
    size: new Map<string, number>()
};

function processLine(line: string) : void {
    const parts = line.split(" ");
    if (parts[0] == "$") {  // no need to do anything with ls
        if (parts[1] == "cd") {
            if (parts[2] == "/") {
                state.current = "/";
            } else if (parts[2] == "..") {
                state.current = state.current.substring(0, state.current.lastIndexOf("/", state.current.length-2) + 1);
            } else {
                state.current += parts[2] + "/";
            }
        }
    } else if (parts[0] != "dir") { // no need to do anything with dir
        state.size.set(state.current + parts[1], Number(parts[0]));
    }
}

function extractDirs(path: string) : string[] {
    const ret: string[] = [];
    for (let i = 0; i != -1; i = path.indexOf("/", i + 1)) {
        ret.push(path.substring(0, i + 1));
    }
    return ret;
}

function run() {
    for (let line of lines) processLine(line); 

    const folders = new Map<string, number>();
    for (let [fullPath, size] of state.size.entries()) {
        for (let path of extractDirs(fullPath)) {
            folders.set(path, (folders.get(path) || 0) + size);
        }
    }
    const sizes = [...folders.values()].sort((a, b) => a - b);
    
    const part1 = sizes.filter(size => size <= 100_000).reduce((a, b) => a + b);
    const part2 = sizes.find(elm => sizes.at(-1)! - elm <= 40_000_000);
    console.debug(part1, part2);
};

run(); // 1477771 3579501
