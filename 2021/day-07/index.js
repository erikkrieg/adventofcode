const fs = require("fs")
const stdinBuffer = fs.readFileSync(0)
const inputs = stdinBuffer.toString().split(',').map(n => parseInt(n, 10))

const crabs = inputs
const totalCrabs = crabs.length

const alignOnMedian = () => {
    const medians = [crabs[totalCrabs/2-1], totalCrabs/2]
    const fuelUsed = [0,0]
    for (let m = 0; m < medians.length; m++) {
        for (let i = 0; i < totalCrabs; i++) {
            fuelUsed[m] += Math.abs(crabs[i] - medians[m])
        }
    }
    console.log(fuelUsed)
}

const alignOnAverage = () => {
    const average = crabs.reduce((a, b) => a+b,0) / totalCrabs
    const averages = [Math.floor(average), Math.ceil(average)]
    const fuelUsed = [0,0]
    for (let a = 0; a < averages.length; a++) {
        for (let i = 0; i < totalCrabs; i++) {
            const delta = Math.abs(crabs[i] - averages[a])
            for (let d = delta; d > 0; d--) {
                fuelUsed[a] += d
            }
        }
    }
    console.log(fuelUsed)
}

alignOnMedian()
alignOnAverage()
