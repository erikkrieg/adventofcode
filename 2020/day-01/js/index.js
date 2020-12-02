const data = require('./data.json')

const dataSet = new Set(data)
const desiredSum = 2020

/**
 * Find a variable number of addends (defaults to two) in a set of unique numbers.
 * 
 * @param {Set} set Set of numbers that will be searched for addends.
 * @param {Number} sum Number that addends must add up to.
 * @param {Number} addendCount The number of unique addends. Greater than 2 results in recursion.
 * @param {Set} ignoreAddends Used to ensure uniqueness of addends during recursion over the set.
 * @return {Array} Array of addends or an empty array if set does not contain requested addends.
 */
const findAddendsInSet = (set, sum, addendCount=2, ignoreAddends) => {
  for (let x of set) {
    const exclude = ignoreAddends || new Set()
    const y = sum - x
    if (exclude.has(x) || x === y) continue
    if (addendCount > 2) {
      exclude.add(x)
      const addends = findAddendsInSet(set, y, addendCount - 1, exclude)
      if (addends.length) {
        return [x, ...addends]
      }
    } else if (set.has(y) && !exclude.has(y)) {
      return [x, y]
    }
  }
  return []
}

const formatAddends = (addends) => {
  if (Array.isArray(addends) && addends.length) {
    console.log(`${addends.join(' + ')} = ${addends.reduce((sum, add) => sum + add)}`)
    console.log(`product = ${addends.reduce((product=1, factor) => product * factor)}`)
  } else {
    console.log('No addends')
  }
}

console.log('Two addends:')
formatAddends(findAddendsInSet(dataSet, desiredSum, 2))

console.log('\nThree addends:')
formatAddends(findAddendsInSet(dataSet, desiredSum, 3))

console.log('\nFour addends:')
formatAddends(findAddendsInSet(dataSet, desiredSum, 4))

console.log('\nFive addends:')
formatAddends(findAddendsInSet(dataSet, desiredSum, 5))
