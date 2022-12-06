const fs = require("fs")
const stdinBuffer = fs.readFileSync(0)
const inputs = stdinBuffer.toString().split('\n')

const productOfGammaAndEpsilonRates = () => {
    const bits = 12
    let gammaRate = '', epsilonRate = '';
    for (let b = 0; b < bits; b++) {
        let zero = 0, one = 0;
        inputs.forEach((bin) => {
            if (bin[b] === '0') zero++
            else one++
        })
        if (zero > one) {
            gammaRate += '0'
            epsilonRate += '1'
        } else {
            gammaRate += '1'
            epsilonRate += '0'
        }
    }
    return parseInt(gammaRate, 2) * parseInt(epsilonRate, 2)
}

const findRatings = (type, numbers) => {
    for (let i = 0; i < 12; i++) {
        const groups = [[], []]
        numbers.forEach((bin) => {
            if (!bin) return
            groups[parseInt(bin[i], 10)].push(bin)
        })
        if (type === 'oxygen') {
            numbers = groups.sort((a, b) => a.length - b.length )[1]
        } else { 
            numbers = groups.sort((a, b) => a.length - b.length )[0]
        }
        if (numbers.length === 1) return numbers[0]
    }
    console.log('didn\'t find a match')
}

const oxy = parseInt(findRatings('oxygen', inputs), 2)
const co2 = parseInt(findRatings('co2', inputs), 2)
console.log(oxy * co2 )
