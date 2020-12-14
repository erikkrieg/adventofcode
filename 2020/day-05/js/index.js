const fs = require('fs')
const data = fs.readFileSync('../data', 'utf-8')

const getSeatId = (boardingPass) => {
  let row = '', column = ''
  for (let i = 0; i < boardingPass.length; i++) {
    const char = boardingPass[i]
    if (i < 7) {
      row += char === 'F' ? '0' : '1'
    } else {
      column += char === 'L' ? '0' : '1'
    }
  }
  return parseInt(row, 2) * 8 + parseInt(column, 2)
}

const sortedSeatIds = data.split('\n').map(getSeatId).sort((a,b) => a - b)
const largestSeatId = sortedSeatIds[sortedSeatIds.length - 1]
const mySeatId = sortedSeatIds.find((id, i, arr) => {
  return (i < arr.length - 1) && (arr[i + 1] !== id + 1)
}) + 1

console.log(largestSeatId, mySeatId)