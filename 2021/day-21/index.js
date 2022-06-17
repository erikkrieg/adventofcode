const fs = require("fs")
const stdinBuffer = fs.readFileSync(0)
const input = stdinBuffer.toString().split('\n')

const Die = () => {
  const max = 100
  return {
    totalRolls: 0,
    roll: function () {
      return (++this.totalRolls % max) || max
    }
  }
}

const play = (playerStarts, D=Die, winningScore=1000) => {
  const die = D()
  const players = playerStarts.map(ps => {
    return {
      player: parseInt(ps.match(/\d+/)[0], 10),
      score: 0,
      pos: parseInt(ps.match(/\d+$/)[0], 10)   
    }
  })
  let i = 0
  while (Math.max(players[0].score, players[1].score) < winningScore) {
    const player = players[i % players.length]
    const rollSum = die.roll() + die.roll() + die.roll()
    player.pos = ((player.pos + rollSum) % 10) || 10
    player.score += player.pos
    i++
  }
  return {
    players,
    totalRolls: die.totalRolls
  }
}

const part1 = ({ players, totalRolls }) => {
  players.sort((a, b) => a.score - b.score)
  return players[0].score * totalRolls
}

const result = play([
  'Player 1 starting position: 4',
  'Player 2 starting position: 8',
])

console.log('part 1 test:', part1(result))
console.log('part 1:', part1(play(input)))


// Part 2 (failed attempt)

const DiracDie = () => {
  return {
    roll: function (players, index) {
      const rolls = [
        [3, 1],
        [4, 3],
        [5, 6],
        [6, 7],
        [7, 6],
        [8, 3],
        [9, 1],
      ]
      return { rolls, total: 3**3 }
    }
  }
}

const quantumPlay = (playerStarts, winningScore=21) => {
  const { roll } = DiracDie()
  // pos and scores and count
  const players = playerStarts.map(ps => {
    return {
      player: parseInt(ps.match(/\d+/)[0], 10),
      scores: new Map([
       [`0:${parseInt(ps.match(/\d+$/)[0], 10)}`, 0],
      ]),
      wins: 0
    }
  })

  do {
    for (let i = 0; i < players.length; i++) {
      const newScores = new Map()
      roll().rolls.forEach(([r, count]) => {
        for (let [key, val] of players[i].scores.entries()) {
          const [oldScore, oldPos] = key.split(':').map(v => parseInt(v, 10))
          if (oldScore >= 21) {
            continue
          }
          const newPos = ((oldPos + r) % 10) || 10
          const newScore = oldScore + newPos
          const newKey = `${newScore}:${newPos}`
          let newCount = (newScores.get(newKey) || 0) + val
          for (let [_,v] of players[(i + 1) % 2].scores.entries()) {
            newCount += count * (v || 1)
          }
          // newCount += count * (3**3 * players[(i + 1) % 2].scores.size)
          if (newScore >= 21) {
            // for (let [k,v] of players[(i + 1) % 2].scores.entries()) {
            //   players[i].wins += newCount * v
            // }
            players[i].wins += newCount
            continue
          }
          newScores.set(newKey, newCount)
        }
      })
      players[i].scores = newScores
      console.log(newScores.size)
    }
  } while (Math.min(players[0].scores.size, players[1].scores.size) > 0)
  console.log(...players)
}

// This solution did not work. There is a working solution using recursion in `part2.js`
// quantumPlay(['Player 1 starting position: 4', 'Player 2 starting position: 8'])