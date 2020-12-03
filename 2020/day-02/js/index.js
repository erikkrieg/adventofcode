const fs = require('fs')
const validateSledRentalPassword = require('./validateSledRentalPassword')
const validateTobogganRentalPassword = require('./validateTobogganRentalPassword')

const passwords = fs.readFileSync('../data', 'utf8').split('\n')
const validSledRentalPasswords = passwords.filter(p => validateSledRentalPassword(...p.split(': ')))
const validTobogganRentalPasswords = passwords.filter(p => validateTobogganRentalPassword(...p.split(': ')))

console.log('Valid sled rental passwords', validSledRentalPasswords.length)
console.log('Valid toboggan rental passwords', validTobogganRentalPasswords.length)
