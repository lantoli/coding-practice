import { readFileSync } from 'fs';

const inputContent = readFileSync("inputs/16_input.txt", 'utf-8');
const sampleContent = readFileSync("inputs/16_sample.txt", 'utf-8');

type Valve = { name: string; rate: number; childNames: string[]; children: Valve[] };
type Player = { current: Valve, visited: Set<string> };
type State = { me: Player, el: Player, open: Set<string>, pressure: number, minute: number, remainRate: number };

function getValves(content: string) {
    const lines = content.split('\n');
    lines.pop(); // Remove last empty line

    const valves: Valve[] = lines.map(line => {
        const [_full, name, rate, childName] = line.match(/Valve (.*) has flow rate=(.*); tunnels? leads? to valves? (.*)/)!;
        return { name, rate: Number(rate), childNames: childName.split(", "), children: [] };
    });
    for (let valve of valves) {
        valve.children = valve.childNames.map(name => valves.find(v => v.name === name)!);
    }
    return valves;
}

function getChildPlayers(player: Player) : Player[] {
    return player.current.children.filter(child => !player.visited.has(child.name)).map(child => 
        ({ current: child, visited: new Set([child.name, ...player.visited])}));
}

const emptyPlayer = (player: Player) => ({ current: player.current, visited: new Set([player.current.name]) });

function run(content: string, minutes: number, withElephant: boolean) {
    const valves: Valve[] = getValves(content);    
    const remainRate = valves.map(valve => valve.rate).reduce((a, b) => a + b);
    const firstValve = valves.find(v => v.name === "AA")!;
    const firstPlayer: Player = { current: firstValve, visited: new Set([firstValve.name]) };
    const states: State[] = [{ me: firstPlayer, el: firstPlayer, open: new Set(), pressure: 0, minute: 1, remainRate }];

    let ret = 0;
    while (states.length > 0) {
        const state = states.pop()!;
        const {me, el, open, pressure, minute, remainRate} = state;
        
        const newMinute = minute + 1;
        const remainingMinutes = minutes - newMinute + 1;
        const optimistic = remainRate * remainingMinutes;
        if (optimistic <= 0 || ret > pressure + optimistic) continue;

        const newMes = getChildPlayers(me);
        const newEls = getChildPlayers(el);
        for (let newMe of newMes) for (let newEl of withElephant ? newEls : [el]) {
            states.push({ me: newMe, el: newEl, open, pressure, minute: newMinute, remainRate });
        }

        const openMe = me.current.rate > 0 && !open.has(me.current.name);
        const openEl = el.current.rate > 0 && !open.has(el.current.name);

        if (openMe) {
            const newPressure = pressure + me.current.rate * remainingMinutes;
            ret = Math.max(ret, newPressure);
            let newState: State = { me: emptyPlayer(me), el, open: new Set([me.current.name, ...open]), 
                pressure: newPressure, minute: newMinute, remainRate: remainRate - me.current.rate };

            if (withElephant && openEl && me.current !== el.current) {
                const newPressure = newState.pressure + el.current.rate * remainingMinutes;
                ret = Math.max(ret, newPressure);
                states.push({...newState, el: emptyPlayer(el), open: new Set([el.current.name, ...newState.open]), 
                    pressure: newPressure, remainRate: newState.remainRate - el.current.rate });
            }

            if (!withElephant) {
                states.push(newState);
            } else for (let newEl of newEls) {
                states.push({...newState, el: newEl });
            }
        }

        if (withElephant && openEl) {
            const newPressure = pressure + el.current.rate * remainingMinutes;
            ret = Math.max(ret, newPressure);
            let newState: State = { me, el: emptyPlayer(el), open: new Set([el.current.name, ...open]), 
                pressure: newPressure, minute: newMinute, remainRate: remainRate - el.current.rate };

            if (openMe && me.current !== el.current) {
                const newPressure = newState.pressure + me.current.rate * remainingMinutes;
                ret = Math.max(ret, newPressure);
                states.push({...newState, me: emptyPlayer(me), open: new Set([me.current.name, ...newState.open]), 
                    pressure: newPressure, remainRate: newState.remainRate - me.current.rate });
                }
            
            for (let newMe of newMes) {
                states.push({...newState, me: newMe });
            }
        }
    }
    console.debug(ret);
};

run(sampleContent, 30, false); // 1651 (sample)
run(sampleContent, 26, true); // 1707 (sample)
run(inputContent, 30, false); // 2029
run(inputContent, 26, true); // 2723

