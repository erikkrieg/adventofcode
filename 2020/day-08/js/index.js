const fs = require('fs')
const data = fs.readFileSync('../data', 'utf-8').split('\n')

const execUniqueInstructions = (instructions, toggleInstruction=-1) => {
  const executedOperationIndexes = new Set()
  let accumulator = 0
  let error
  
  let i = 0;
  while (!error && i < instructions.length) {
    error = executedOperationIndexes.has(i)
    executedOperationIndexes.add(i)
    let [op, arg] = instructions[i].split(' ')
    if (i === toggleInstruction) {
      op = op === 'jmp' ? 'nop' : 'jmp'
    }
    switch (op) {
      case 'acc':
        accumulator += parseInt(arg, 10)
        i++
        break
      case 'jmp':
        i += parseInt(arg, 10)
        break
      default:
        i++
    }
    error = executedOperationIndexes.has(i)
  }
  return { accumulator, executedOperationIndexes, error }
}

const execCorruptedInstructions = (instructions, subset) => {
  const scopedInstructions = subset || instructions
  for (const i of scopedInstructions) {
    const [op, arg] = instructions[i].split(' ')
    if (op !== 'acc') {
      const { error, accumulator } = execUniqueInstructions(instructions, i)
      if (!error) return { accumulator }
    }
  }
  throw new Error('Program is too corrupted!')
}

const { accumulator, executedOperationIndexes } = execUniqueInstructions(data)

console.log('execUniqueInstructions', accumulator)
console.log('execCorruptedInstructions', execCorruptedInstructions(data, executedOperationIndexes).accumulator)


