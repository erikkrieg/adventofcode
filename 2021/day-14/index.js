const fs = require("fs")
const stdinBuffer = fs.readFileSync(0)
const [template, unformatedInserts] = stdinBuffer.toString().split('\n\n')
const inserts = unformatedInserts.split('\n').reduce((m, i) => m.set(...i.split(' -> ')), new Map())

const insertion = (formula, ins) => {
  let newFormula = formula[0]
  for (let i = 1; i < formula.length; i++) {
    const key = formula.slice(i-1, i+1)
    newFormula += `${ins.get(key)}${key[1]}`
  }
  return newFormula
}

const count = (formula) => {
  const countMap = new Map()
  for (let i = 0; i < latestFormula.length; i++) {
    const k = latestFormula[i]
    countMap.set(k, (countMap.get(k) || 0) + 1)
  }
  
  const sortedCount = [...countMap].sort((a,b) => a[1] - b[1])
  return sortedCount
}

let latestFormula = template
for (let steps = 0; steps < 10; steps++) {
  latestFormula = insertion(latestFormula, inserts)
  console.log(
    steps+1,
    count(latestFormula)
  )
}

/*
This can determine the number of elements in a "formula"

const est = (len, step) => {
  step - 1
  return (len - 1) * (2**(step)) + 1
}

3298534883329
3,298,534,883,329
*/