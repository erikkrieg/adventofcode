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

const numberInRange = (n, gte, lte) => {
  const num = Number(n)
  return num >= gte && num <= lte
}
const eyeColours = new Set(['amb', 'blu', 'brn', 'gry', 'grn', 'hzl', 'oth'])
const pidRegex = /^[0-9]{9}$/
const validator = {
  // byr (Birth Year) - four digits; at least 1920 and at most 2002.
  byr: (n) => numberInRange(n, 1920, 2002),
  // iyr (Issue Year) - four digits; at least 2010 and at most 2020.
  iyr: (n) => numberInRange(n, 2010, 2020),
  // eyr (Expiration Year) - four digits; at least 2020 and at most 2030.
  eyr: (n) => numberInRange(n, 2020, 2030),
  // hgt (Height) - a number followed by either cm or in:
  // If cm, the number must be at least 150 and at most 193.
  // If in, the number must be at least 59 and at most 76.
  hgt: (n) => {
    if (n.includes('cm')) {
      return numberInRange(parseInt(n, 10), 150, 193)
    } else if (n.includes('in')) {
      return numberInRange(parseInt(n, 10), 59, 76)
    } else {
      return false
    }
  },
  // hcl (Hair Color) - a # followed by exactly six characters 0-9 or a-f.
  hcl: (n) => {
    const [_, hex] = n.split('#')
    return parseInt(hex, 16).toString(16).padStart(6, '0') === hex
  },
  // ecl (Eye Color) - exactly one of: amb blu brn gry grn hzl oth.
  ecl: (s) => {
    return eyeColours.has(s)
  },
  // pid (Passport ID) - a nine-digit number, including leading zeroes.
  pid: (s) => {
    return pidRegex.test(s)
  }
}

const countValidPassportsV2 = (passports, validator) => {
  const passportFieldRegex = /(\w+):(\S+)/g
  let validPassports = 0
  passports.forEach((passportStr, i) => {
    let validFields = 0
    for (const [_, key, value] of passportStr.matchAll(passportFieldRegex)) {
      if (key === 'cid') continue
      if (!validator[key](value)) return false
      validFields++
    }
    validPassports += validFields === 7 
  })
  return validPassports
}

const passports = data.split('\n\n')
console.log(
  'Valid passports:',
  countValidPassports(passports)
)

console.log(
  'Valid passports (v2):',
  countValidPassportsV2(passports, validator)
)
