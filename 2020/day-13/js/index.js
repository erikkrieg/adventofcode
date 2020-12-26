const fs = require('fs')
const [departTimeRaw, busIdsRaw] = fs.readFileSync('../data', 'utf-8').split('\n')
const departTime = Number(departTimeRaw)

const findBusWithShortestWait = (departTime, busIds) => {
  const idealBus = { wait: Infinity, id: null }
  busIds = busIds.match(/\d+/g)
  busIds.forEach((idStr) => {
    const id = Number(idStr)
    const wait = id - (departTime % id)
    if (wait < idealBus.wait) {
      idealBus.wait = wait
      idealBus.id = id
    }
  })
  return idealBus
}

const { wait, id } = findBusWithShortestWait(departTime, busIdsRaw)
console.log(wait * id)

// Part 2

const multiplicativeInverses = (p, m) => {
  let a = p, b = m
  let t = 0, s = 1
  while (b > 0) {
    const quotient = Math.floor(a / b)
    const remainder = a % b
    const temp = t
    a = b
    b = remainder
    t = s - quotient * t
    s = temp
  }
  return s < 0 ? s + m : s
}

const crt = ({ m, a }) => {
  const product = m.reduce((p, f) => p * f, 1)
  let sum = 0
  for (let i = 0; i < m.length; i++) {
    const p = product / m[i]
    sum += a[i] * p * multiplicativeInverses(p, m[i])
  }

  return sum % product
}

const part2 = (busIds) => {
  let cur = ''
  let offset = 0
  const buses = []
  for (let i = 0; i <= busIds.length; i++) {
    if(/\d+/.test(busIds[i])) {
      cur += busIds[i]
    } else {
      if (cur) {
        buses.push({ id: Number(cur), offset })
        offset++
        cur = ''
      }
      if (busIds[i] === 'x') {
        offset++
      }
    }
  }

  return crt({
    m: buses.map(({ id }) => id),
    a: buses.map(({ id, offset }) => (id - offset) % id),
  })
}

// For testing smaller sets:
// console.log(part2('67,x,7,59,61'))
// console.log(part2('1789,37,47,1889'))

console.log(part2(busIdsRaw))
