const fs = require('fs')
const data = fs.readFileSync('../data', 'utf-8').split('\n')
const joltageRatings = data.map(Number).sort((a,b)=>a-b)

// Part 1 - Requires sorted array of adapter joltage ratings
const countJoltageRatingDifferenceDistribution = (ratings) => {
  let lastJoltage = 0
  return ratings.reduce((diffs, joltage) => {
    const diff = joltage - lastJoltage
    diffs[diff] = (diffs[diff] || 0) + 1
    lastJoltage = joltage
    return diffs
  }, { '3': 1 })
}

/**
 * Part 2 - Requires sorted array of adapter joltage ratings
 * 
 * There are 3 distinct patterns that can be found in a collection of adapters:
 * 1. An adapter that is required to complete a combination.
 * 2. An adapter that is optional to complete a combination.
 * 3. A group of optional adapters where at least one adapter from the group is required.
 * 
 */
const describeAdapterRatings = (ratings) => {
  const ratingsLen = ratings.length
  const optionalGroups = []
  let required = 0
  let optional = 0
  let lastRating = 0
  let lastReq = 0
  let consecutiveOptional = 0
  for (let i = 0; i < ratingsLen; i++) {
    let options = 0
    for (let x = i; x < ratingsLen && x < i + 3; x++) {
      if (ratings[x] <= (lastRating + 3)) options++
    }
    if (options === 1) {
      const lastReqDiff = ratings[i] - lastReq
      if (consecutiveOptional > 1 && lastReqDiff > 3) {
        optionalGroups.push(consecutiveOptional)
        optional -= consecutiveOptional
      }
      consecutiveOptional = 0
      lastReq = ratings[i]
      required++
    } else {
      optional++
      consecutiveOptional++
    }
    lastRating = ratings[i]
  }
  return {
    required,
    optional,
    optionalGroups
  }
}

const countAdapterCombinations = (ratings) => {
  const { optional, optionalGroups } = describeAdapterRatings(ratings)
  const optionalCombos = 2 ** optional
  const optionalSetCombos = optionalGroups.reduce((product, factor) => product * (2 ** factor - 1), 1)

  return optionalCombos * optionalSetCombos
}

console.log(countAdapterCombinations(joltageRatings))
