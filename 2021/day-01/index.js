const fs = require("fs")
const stdinBuffer = fs.readFileSync(0)
const inputs = stdinBuffer.toString().split('\n').map(n => parseInt(n, 10))

const countIncrease = () => {
    let increases = 0
    for (let i = 1; i < inputs.length; i++) {
        increases += inputs[i] > inputs[i-1]
    }
    
    console.log(increases)
}

const countSlidingWindowIncreases = () => {
    const windowRadius = 1
    let increases = 0
    let lastWindowSum
    for (let i = windowRadius; i + windowRadius < inputs.length; i++) {
        const curWindowSum = inputs[i - windowRadius] + inputs[i] + inputs[i + windowRadius]
        if (lastWindowSum) {
            increases += curWindowSum > lastWindowSum
        }
        lastWindowSum = curWindowSum
    }
    console.log(increases)
}
countIncrease()
countSlidingWindowIncreases()
