const fs = require('fs')
const data = fs.readFileSync('../data', 'utf-8')

const countValidPassports = (passports) => {
  const optionalField = 'cid:'
  let validCount = 0
  passports.forEach((passport) => {
    const passportFields = new Set(passport.match(/(\w{3}):/g))
    if (
      passportFields.size === 8 ||
      (passportFields.size === 7 && !passportFields.has(optionalField))
    ) {
      validCount++
    }
  })
  return validCount
}

console.log(
  'Valid passports:',
  countValidPassports(data.split('\n\n'))
)
