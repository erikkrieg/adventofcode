const fs = require('fs')
const data = fs.readFileSync('../data', 'utf-8').split('\n')

const getCoords = (rawSeatsStr) => {
  const SEAT = 'L', FLOOR = '.'
  const seats = new Set(), floor = new Set()
  rawSeatsStr.forEach((row, y) => {
    for (let x = 0; x < row.length; x++) {
      const coord = `${x},${y}`
      if (row[x] === SEAT) {
        seats.add(coord)
      } else if (row[x] === FLOOR) {
        floor.add(coord)
      }
    }
  })
  return { seats, floor }
}

const countAdjacentOccupiedSeats = (seat, seats, floor, occupiedSeats) => {
  const [x, y] = seat.split(',').map(Number)
  const diffs = [
    [1,0], [-1,0],
    [0, 1], [1,1], [-1,1],
    [0, -1], [1,-1], [-1,-1],
  ]
  let adjacentSeats = 0
  diffs.forEach(([dx, dy]) => {
    let seat = `${x + dx},${y + dy}`, i = 2
    while (floor.has(seat)) {
      seat = `${x + dx * i},${y + dy * i}`
      i++
    }
    adjacentSeats += occupiedSeats.has(seat)
  })
  return adjacentSeats
}

const getSeatDiff = (seats, floor, occupiedSeats) => {
  const add = new Set(), remove = new Set()
  for (const seatCoord of seats) {
    const adjOccupied = countAdjacentOccupiedSeats(seatCoord, seats, floor, occupiedSeats)
    if (occupiedSeats.has(seatCoord)) {
      if (adjOccupied > 4) remove.add(seatCoord)
    } else {
      if (adjOccupied === 0) add.add(seatCoord)
    }
  }
  return { add, remove }
}

const musicalChairs = (seatGrid) => {
  const { seats, floor } = getCoords(seatGrid)
  let occupiedSeats = new Set()
  let hasChanges = true
  let rounds = 0
  while (hasChanges) {
    const { add, remove } = getSeatDiff(seats, floor, occupiedSeats)
    hasChanges = remove.size + add.size > 0
    occupiedSeats = new Set(
      [...occupiedSeats]
        .filter((s) => !remove.has(s))
        .concat([...add])
    )
    rounds++
  }
  console.log('Round', rounds)
  console.log('- Occupied:', occupiedSeats.size)
}

musicalChairs(data)
