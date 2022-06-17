const fs = require("fs")
const stdinBuffer = fs.readFileSync(0)
const inputs = stdinBuffer.toString().split('\n').map(x => x.split(' '))

const productOfMoves = (moves) => {
    let h = 0, 
        d = 0;
    
    for (let i = 0; i < inputs.length; i++) {
        const [dir, val] = inputs[i]
        if (dir === 'forward') {
            h += parseInt(val, 10)
        } else if (dir === 'down') {
            d += parseInt(val, 10)
        } else if (dir === 'up') {
            d -= parseInt(val, 10)
        }
    }
    return h * d
}

// console.log(productOfMoves(inputs))

const productOfAimedMoves = (moves) => {
    let h = 0, 
        d = 0,
        a = 0;
    for (let i = 0; i < inputs.length; i++) {
        const dir = inputs[i][0]
        const val = parseInt(inputs[i][1], 10)
        if (dir === 'forward') {
            h += val
            d += val * a
        } else if (dir === 'down') {
            a += val
        } else if (dir === 'up') {
            a -= val
        }
    }
    return h * d
}

console.log(productOfAimedMoves(inputs))

