const fs = require("fs")
const stdinBuffer = fs.readFileSync(0)
const inputs = stdinBuffer.toString().split('\n\n').map(i => i.split('\n'))
const [dotsArr, folds] = inputs
const dots = new Set([...dotsArr])

folds.forEach((fold) => {
  const [axis, posStr] = fold.match(/[xy]{1}\=\d+/)[0].split('=')
  const pos = parseInt(posStr, 10)
  dots.forEach((d) => {
    const [x,y] = d.split(',').map(c => parseInt(c, 10))
    if (axis === 'y' && pos < y) {
      dots.add(`${x},${pos - (y - pos)}`)
      dots.delete(d)
    } else if (axis === 'x' && pos < x) {
      dots.add(`${pos - (x - pos)},${y}`)
      dots.delete(d)
    }
  })
  // Part 1 requires number of dots after the first fold.
  // console.log(dots.size)
})

const out = [...dots].map(d => d.split(',').map(dd => parseInt(dd, 10)))
const maxX = out.sort((a,b) => a[0] - b[0])[out.length-1][0]
const maxY = out.sort((a,b) => a[1] - b[1])[out.length-1][1]

for (let y = 0; y <= maxY; y++) {
  const row = []
  for (let x = 0; x <= maxX; x++) {
    row.push(dots.has(`${x},${y}`) ? ' #' : '  ')
  }
  console.log(row.join(''))
}