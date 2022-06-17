const fs = require("fs")
const stdinBuffer = fs.readFileSync(0)
const [numbersStr, ...boardStrs] = (stdinBuffer.toString() + '\n').split('\n\n')

const numbers = numbersStr.split(',').map(s => parseInt(s, 10))
const boards = boardStrs.map(b => b.split(/\W/).filter(b => !!b))

const findWinner = () => {
    const winningScores = []
    for (let n = 0; n < numbers.length; n++) {
        for (let b = 0; b < boards.length; b++) {
            for (let i = 0; i < boards[b].length; i++) {
                if (boards[b][i] == numbers[n]) {
                    boards[b][i] = null
                    const row = Math.floor(i / 5) * 5
                    const col = i % 5
                    let rowMatches = 0, colMatches = 0;
                    for (let ii = 0; ii < 5; ii++) {
                        rowMatches += boards[b][ii + row] === null
                        colMatches += boards[b][ii * 5 + col] === null
                    }
                    if ([rowMatches, colMatches].includes(5)) {
                        const score = getScore(numbers[n], boards[b])
                        winningScores.push(score)
                        boards[b] = []
                        break;
                    }
                }
            }
        }
    }
    return winningScores
}

const getScore = (factor, board) => {
    const boardSum = board.reduce((sum, num) => sum - -num, 0)
    const score = factor * boardSum
    return score
}

const winners = findWinner()
console.log('First winning score:', winners[0])
console.log('Last winning score:', winners[winners.length - 1])