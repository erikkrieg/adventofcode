const fs = require("fs")
const stdinBuffer = fs.readFileSync(0)
const inputs = stdinBuffer.toString().split(',')

// Original part 1 solution: O(n^2)
//
// const fishTimers = [inputs.map(n => parseInt(n, 10))]
// const days = 256
// const spawnTime = 6
// const firstSpawnTimeOffset = 2
//
// console.log(fishTimers.length)
// for (let d = days; d > 0; d--) {
//     fishTimers.push([])
//     for (let f = fishTimers.length - 2; f >= 0; f--) {
//         for (let ff = fishTimers[f].length -1; ff >= 0; ff--) {
//             if (fishTimers[f][ff] === 0) {
//                 fishTimers[f][ff] = spawnTime
//                 fishTimers[fishTimers.length - 1].push(spawnTime + firstSpawnTimeOffset)
//             } else {
//                 fishTimers[f][ff] -= 1
//             }
//         }
//     }
// }
// const population = fishTimers.reduce((sum, cur) => sum + cur.length, 0)
// console.log(population)

const fishTimers = inputs.map(n => parseInt(n, 10))
const simulatePopulation = (initialTimers, periodInDays) => {
    const maxSpawnPeriod = 9
    const minSpawnPeriod = 7
    const timers = Array(maxSpawnPeriod).fill(0)
    for (let i = 0; i < initialTimers.length; i++) {
        timers[initialTimers[i]]++
    }
    for (let d = periodInDays; d > 0; d--) {
        let newFish = timers[0]
        for (let t = 1; t < timers.length; t++) {
            timers[t - 1] = timers[t]
        }
        timers[minSpawnPeriod - 1] += newFish
        timers[maxSpawnPeriod - 1] = newFish
    }
    return timers.reduce((sum, cur) => sum + cur, 0)
}

console.log(simulatePopulation(fishTimers, 80))
console.log(simulatePopulation(fishTimers, 256))