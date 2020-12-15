const fs = require('fs')
const data = fs.readFileSync('../data', 'utf-8')

const bagNameReg = /\w+ \w+ bag/g
const luggageRules = data.split('\n').reduce((rules, rule) => {
  const bagNames = rule.match(bagNameReg)
  if (bagNames.length === 2 && bagNames[1] === 'no other bag') return rules
  for (let i = bagNames.length - 1; i > 0; i--) {
    rules[bagNames[i]] = rules[bagNames[i]] || []
    rules[bagNames[i]].push(bagNames[0])
  }
  return rules
}, {})

const countValidOuterBags = (bagName, validBagSet = new Set()) => {
  const outerBags = luggageRules[bagName] || []
  outerBags.forEach((b) => {
    if (!validBagSet.has(b)) {
      validBagSet.add(b)
      countValidOuterBags(b, validBagSet)
    }
  })
  return validBagSet.size
}

console.log(countValidOuterBags('shiny gold bag'))
