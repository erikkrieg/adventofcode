const fs = require('fs')
const data = fs.readFileSync('../data', 'utf-8').split('\n')

const part1 = (instructions) => {
  const dirs = 'NESW'
  const mod = {
    N: { x: 0, y: 1 },
    S: { x: 0, y: -1 },
    E: { x: 1, y: 0 },
    W: { x: -1, y: 0 },
  }
  const pos = { x: 0, y: 0 }
  let facing = 1
  instructions.forEach(([ins, ...val]) => {
    val = Number(val.join(''))

    if ('LR'.includes(ins)) {
      facing = (val / 90 * (ins === 'L' ? -1 : 1) + facing) % dirs.length
      if (facing < 0) {
        facing = (facing + dirs.length) % dirs.length
      }
    } else {
      curDir = mod[mod[ins] ? ins : dirs[facing]]
      pos.x += curDir.x * val
      pos.y += curDir.y * val
    }
  })
  return Math.abs(pos.x) + Math.abs(pos.y)
}

// Only rotates points by degrees that are divisible by 90.
const rotate = ({x, y}, rot, dir='R') => {
  const aMod = dir === 'R' ? 1 : -1
  const bMod = dir === 'R' ? -1 : 1
  for (let i = rot / 90; i > 0; i--) {
    const a = y * aMod
    const b = x * bMod
    x = a, y = b
  }
  return { x, y }
}

const part2 = (instructions) => {
  const mod = {
    N: { x: 0, y: 1 },
    S: { x: 0, y: -1 },
    E: { x: 1, y: 0 },
    W: { x: -1, y: 0 },
  }

  const shipPos = { x: 0, y: 0 }
  const waypointPos = { x: 10, y: 1 }

  instructions.forEach(([ins, ...val]) => {
    val = Number(val.join(''))
    if ('LR'.includes(ins)) {
      Object.assign(waypointPos, rotate(waypointPos, val, ins))
    } else if (ins === 'F') {
      shipPos.x += waypointPos.x * val
      shipPos.y += waypointPos.y * val
    } else {
      waypointPos.x += mod[ins].x * val
      waypointPos.y += mod[ins].y * val
    }
  })
  return { shipPos, waypointPos }
}

console.log('Part 1:', part1(data))
console.log('Part 2:', part2(data))
