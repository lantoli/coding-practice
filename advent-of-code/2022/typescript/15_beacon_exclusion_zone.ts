import { readFileSync } from 'fs';

const inputContent = readFileSync("inputs/15_input.txt", 'utf-8');
const sampleContent = readFileSync("inputs/15_sample.txt", 'utf-8');

type Sensor = { y: number; x: number; range: number; };

function getSensorsAndBaconCount(content: string, y?: number) : [Sensor[], number] {
    const lines = content.split('\n');
    lines.pop(); // Remove last empty line

    const sensors: Sensor[] = [];
    const bacons = new Set<number>();
    for (const line of lines) {
        const [_full, xsensor, ysensor, xbacon, ybacon] = line.match(/Sensor at x=(.*), y=(.*): closest beacon is at x=(.*), y=(.*)/)!.map(Number);
        sensors.push({y: ysensor, x: xsensor, range: Math.abs(ybacon - ysensor) + Math.abs(xbacon - xsensor)});
        if (ybacon === y) bacons.add(xbacon);
    }
    return [sensors, bacons.size];
}

function getIntervals(sensors: Sensor[], y: number) : [number, number][] {
    const intervals: [number, number][] = [];
    for (let sensor of sensors) {
        const ydist = sensor.range - Math.abs(sensor.y - y);
        if (ydist >= 0) {
            intervals.push([sensor.x - ydist, sensor.x + ydist]);    
        }
    }    
    intervals.sort((a, b) => a[0] - b[0]);
    return intervals;
}

function run1(content: string, y: number) {
    const [sensors, baconCount] = getSensorsAndBaconCount(content, y);
    const intervals = getIntervals(sensors, y);
    let current = intervals[0][0] - 1;
    let count = 0;
    for (let interval of intervals) {
        const left = Math.max(current + 1, interval[0]);
        const right = interval[1];
        if (left <= right) {
            count += right - left + 1;
        }
        current = Math.max(current, interval[1]);
    }
    console.debug(count - baconCount);
};

function run2(content: string, tunning: number, max: number) {
    const [sensors, _] = getSensorsAndBaconCount(content);
    for (let y = 0; y <= max; y++) {
        const intervals = getIntervals(sensors, y);
        let current = 0;
        for (let interval of intervals) {
            if (current < interval[0]) {
                console.debug(current * tunning + y);
                return;
            }
            current = Math.max(current, interval[1] + 1);
            if (current > max) break;
        }
    }
}

run1(sampleContent, 10); // 26 (sample)
run2(sampleContent, 4_000_000, 20); // 56000011 (sample)
run1(inputContent, 2_000_000); // 5125700
run2(inputContent, 4_000_000, 4_000_000); // 11379394658764
