const fs = require("fs")
const stdinBuffer = fs.readFileSync(0)
const inputs = stdinBuffer.toString().split('\n')
const lines = inputs.reduce((ls, i) => {
    if (!i) return ls
    const line = i.split(' -> ').map(l => l.split(',').map(l => parseInt(l, 10)))
    ls.push(line)
    return ls
}, [])

const countIntersectingPoints = (lines, skipDiagonal=true) => {
    const points = new Map()
    for (let l = 0; l < lines.length; l++) {
        let [[x1, y1], [x2, y2]] = lines[l]
        if (skipDiagonal && x1 !== x2 && y1 !== y2) continue
        while (x1 !== x2 || y1 !== y2) {
            const key = `${x1},${y1}`
            points.set(key, (points.get(key) || 0) + 1)
            x1 += Math.sign(x2 - x1)
            y1 += Math.sign(y2 - y1)
        }
        // The loop won't track the last point based on current logic.
        const key = `${x1},${y1}`
        points.set(key, (points.get(key) || 0) + 1)
    }
    let overlappingPoints = 0
    points.forEach(p => overlappingPoints += p > 1)
    return overlappingPoints
}

console.log('Horizontal and vertical line intersecting points:', countIntersectingPoints(lines))
console.log('Horizontal, vertical and diagonal line intersecting points:',countIntersectingPoints(lines, false))
