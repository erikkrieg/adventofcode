const fs = require("fs")
const stdinBuffer = fs.readFileSync(0)
const inputs = stdinBuffer.toString().split('\n').map(i => i.split('').map(ii => parseInt(ii, 10)))

const log = (octopi) => {
  let out = ''
  octopi.forEach(row => out += row.join('') + '\n')
  console.log(out)
}
const step = (octopi) => {
  const canFlash = new Set()

  // increment
  for (let y = 0; y < octopi.length; y++) {
    for (let x = 0; x < octopi[y].length; x++) {
      octopi[y][x] += 1
      if (octopi[y][x] > 9) {
        canFlash.add(`${x},${y}`)
      }
    }
  }

  // Find all that can flash
  canFlash.forEach((coord) => {
    const [x,y] = coord.split(',').map(c => parseInt(c, 10))
    const adj = [
      [x - 1, y - 1], [x, y - 1], [x + 1, y - 1],
      [x - 1, y], [x + 1, y],
      [x - 1, y + 1], [x, y + 1], [x + 1, y + 1],
    ]
    adj.forEach(([xx,yy]) => {
      const adjCoord = `${xx},${yy}`
      let oct = octopi[yy]?.[xx] || NaN
      if (isNaN(oct) || canFlash.has(adjCoord)) return
      octopi[yy][xx] = ++oct
      if (oct > 9) canFlash.add(adjCoord)
    })
  })

  // Execute flashing
  canFlash.forEach((coord) => {
    const [x,y] = coord.split(',')
    octopi[y][x] = 0
  })

  return {
    octopi,
    flashes: canFlash.size
  }
}

let flashCount = 0
for (let s = 0; s < 1000; s++) {
  const { flashes } = step(inputs)
  flashCount += flashes

  if (flashes === 100) {
    console.log('All flashed on step:', s+1)
    break
  }
}

console.log(flashCount)
