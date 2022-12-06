const fs = require("fs")
const stdinBuffer = fs.readFileSync(0)
const inputs = stdinBuffer.toString().split('\n')
inputs.pop()

const key = {
    '2': 1,
    '4': 4,
    '3': 7,
    '7': 8
}
const countUnique = () => {
    
    let count = 0
    for (let i = 0; i < inputs.length; i++) {
        const signals = inputs[i].split(' | ')[1].split(' ')
        count += signals.reduce((acc, s) => {
            if (key[s.length]) return acc + 1
            return acc
        }, 0)
    }
    console.log(count)
}



const contains = (arr, tar) => tar.split('').every(t => arr.includes(t))
const countMatches = (arr, tar) => tar.split('').filter(t => arr.includes(t)).length

let sum = 0
for (let i = 0; i < inputs.length; i++) {
    const signals = inputs[i].split(' | ').map(s => s.split(' '))
    const randomKey = new Map()
    const codedSignals = []
    signals[0].forEach(s => {
        const sorted = s.split('').sort().join('')
        if (key[sorted.length]) {
            randomKey.set(key[sorted.length], sorted)
        } else {
            codedSignals.push(sorted)
        }
    })
    codedSignals.forEach(s => {
        if (s.length === 5) {
            if (contains(s, randomKey.get(1))) {
                randomKey.set(3, s)
            } else if (countMatches(s, randomKey.get(4)) == 2) {
                randomKey.set(2, s)
            } else if (countMatches(s, randomKey.get(4)) == 3) {
                randomKey.set(5, s)
            }
        } else {
            if (!contains(s, randomKey.get(1))) {
                randomKey.set(6, s)
            } else if (!contains(s, randomKey.get(4))) {
                randomKey.set(0, s)
            } else if (contains(s, randomKey.get(4))) {
                randomKey.set(9, s)
            }
        }
    })
    
    let segment = ''
    signals[1].forEach(s => {
        const sorted = s.split('').sort().join('')
        for (let l = 0; l < randomKey.size; l++) {
            if (sorted === randomKey.get(l)) {
                segment += l
                break
            }
        }
    })
    sum += parseInt(segment, 10)
}
console.log(sum)
