const fs = require('fs')
const data = fs.readFileSync('../data', 'utf-8').split('\n\n')

const countOfYesAnswers = data.reduce((sum, d) => {
  const group = new Set(d.split('\n').join(''))
  return sum + group.size
}, 0)

const countOfUnanimousYesAnswers = data.reduce((sum, d) => {
  const answers = d.split('\n')
  const answersString = answers.join('')
  const counter = {}
  let addend = 0
  for (let i = 0; i < answersString.length; i++) {
    counter[answersString[i]] = counter[answersString[i]] || 0
    counter[answersString[i]]++
    if (counter[answersString[i]] === answers.length) {
      addend++
    }
  }
  return sum + addend
}, 0)

console.log(countOfYesAnswers, countOfUnanimousYesAnswers)
