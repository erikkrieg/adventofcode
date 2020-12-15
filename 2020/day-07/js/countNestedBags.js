const fs = require('fs')
const data = fs.readFileSync('../data', 'utf-8')

const bagNameReg = /\w+ \w+ bag/
const innerBagsReg = /\d \w+ \w+ bag/g
const luggageRules = data.split('\n').reduce((rules, rule) => {
  const outerBag = rule.match(bagNameReg)[0]
  const innerBags = rule.match(innerBagsReg) || []
  rules[outerBag] = innerBags.reduce((acc, cur) => {
    acc[cur.match(bagNameReg)[0]] = parseInt(cur, 10)
    return acc
  }, {})
  return rules
}, {})

const countBags = (bagName, count=1) => {
  const rules = luggageRules[bagName]
  return Object.getOwnPropertyNames(rules).reduce((sum, rule) => {
    return sum + (countBags(rule) * rules[rule])
  }, count)
}

console.log(
  // subtract 1 because we don't want to include the top level bag.
  countBags('shiny gold bag') - 1
)
