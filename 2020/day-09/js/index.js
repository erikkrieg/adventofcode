const fs = require('fs')
const data = fs.readFileSync('../data', 'utf-8').split('\n')

const findInvalidXmasNum = (numbers) => {
  for (let i = 25; i < numbers.length; i++) {
    const addends = new Set(numbers.slice(i - 25, i))
    const sum = parseInt(numbers[i], 10)
    let invalid = true
    for (let x of addends) {
      const y = (sum - parseInt(x, 10)).toString()
      if (addends.has(y) && x !== y) {
        invalid = false
        break
      }
    }
    if (invalid) return sum
  }
}

const findXmasWeakness = (num, addends) => {
  let sum = 0, start = 0, end = 0, found = false
  for (let i = 0; i < addends.length; i++) {
    if (num === sum) {
      found = true
      break
    }

    const iNum = parseInt(addends[i], 10)
    end = i
    sum += iNum

    if (iNum === num || iNum > num) {
      start = i + 1
      sum = 0
    } else if (sum > num) {
      for (let x = start; x < end; x++) {
        const xNum = parseInt(addends[x], 10)
        start = x + 1
        sum -= xNum
        if (sum === num || sum < num) {
          break
        }
      }
    }
  }
  if (found) {
    const range = addends.slice(start, end).map(Number).sort((a,b) => a-b)
    return range[0] + range[range.length - 1]
  }
  return null
}

const invalid = findInvalidXmasNum(data)
const weakness = findXmasWeakness(parseInt(invalid, 10), data)
console.log('Invalid:', invalid)
console.log('Weakness:', weakness)
