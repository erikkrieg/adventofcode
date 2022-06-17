// let p1 = {
//   games: 1,
// }
// let p2 = {
//   games: 1,
// }

// const roll = () => {
//   return [
//     [3, 1],
//     [4, 3],
//     [5, 6],
//     [6, 7],
//     [7, 6],
//     [8, 3],
//     [9, 1],
//   ]
// }

// const playRound = (games=1, iterations=21) => {
//   // will need to hold more state
//   const players = [
//     new Map([[4, [0]]]),
//   ]
//   const rolls = roll()
//   while (iterations>0) {
//     players.map((p) => {
//       const games = new Map()
//       for (let [pos, scores] of p.entries()) {
//         rolls.forEach(([r, count]) => {
//           const newPos = pos + r

//           // const cur = games.get(newPos) || [1]
//           // const newCount = cur.map((c) => c * count)
//           // games.set(newPos, newCount)
//         })

//       }
//     })
//     iterations--
//   }
//   return games
// }

// console.log(playRound(1, 2))

/* 

rounds = [
  {
    player: 1
    positions: {
      scores: {
        [score]: count
      }
    }
  },
  {
    player: 2
    positions: {
      scores: {}
    }
  }
]

rounds = [
  {

  }
]
*/

const fs = require("fs")
const stdinBuffer = fs.readFileSync(0)
const input = stdinBuffer.toString().split('\n')

const startGame = (playerStarts) => {
  const rolls = [
    [3, 1],
    [4, 3],
    [5, 6],
    [6, 7],
    [7, 6],
    [8, 3],
    [9, 1],
  ]
  const totalWins = {
    p1: 0,
    p2: 0,
  }
  const initialPlayerState = playerStarts.map(ps => {
    const id = 'p' + parseInt(ps.match(/\d+/)[0], 10)
    const score = 0
    totalWins[id] = score
    return {
      id,
      score,
      pos: parseInt(ps.match(/\d+$/)[0], 10),
      count: 1,
    }
  })
  const play = (players=initialPlayerState) => {
    const boardSize = 10
    const player = players.shift()
    rolls.forEach(([roll, permutations]) => {
      const pos = ((player.pos + roll) % boardSize) || boardSize
      const score = player.score + pos
      const count = (players[0].count * permutations)
      if (score >= 21) {
        totalWins[player.id] += count
      } else {
        play([players[0], {
          id: player.id,
          pos,
          score,
          count
        }])
      }
    })
  }
  return {
    play,
    getWins() {
      return totalWins
    },
  }
}

// Test input
// const game = startGame([
//   "Player 1 starting position: 4",
//   "Player 2 starting position: 8",
// ])
const game = startGame(input)
game.play()
console.log(game.getWins())
