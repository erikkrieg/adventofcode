const fs = require("fs")
const stdinBuffer = fs.readFileSync(0)
const input = stdinBuffer.toString().split('\n').map(i => i.split(''))

const tileGraph = (graph, times=5) => {
  const tiledGraph = []
  for (let y = 0; y < times; y++) {
    for (let i = 0; i < graph.length; i++) {
      const row = []
      for (let x = 0; x < 5; x++) {
        row.push(...graph[i].map(g => {
          let v = parseInt(g, 10) + x + y
          return v % 9 || v
        }))
      }
      tiledGraph.push(row)
    }
  }
  return tiledGraph
}

const findShortestPath = (graph) => {
  const maxLen = graph.length
  const destination = `${maxLen-1},${maxLen-1}`
  const edgeNodes = new Map([['0,0', { distance: 0, nodes: ['0,1', '1,0']}]])
  
  let steps = 0
  while (edgeNodes.size > 0) {
    steps++
    for (const [node, {distance, nodes}] of edgeNodes.entries()) {
      const remainingNodes = []
      // console.log({node, distance, steps})
      for (let i = 0; i < nodes.length; i++) {
        const [x,y] = nodes[i].split(',').map(c => parseInt(c,10))
        const nextDist = distance + parseInt(graph[y][x], 10)
        // console.log('  ', {n:nodes[i], nextDist}, parseInt(graph[y][x], 10))
        if (nextDist <= steps) {
          if (nodes[i] === destination) return nextDist
          if (edgeNodes.has(nodes[i])) continue
          const nextNodes = [[x + 1, y],[x, y + 1],[x - 1, y],[x, y - 1]]
            .filter(([xx,yy]) => {
              return graph[yy]?.[xx]
            })
            .map(([xx,yy]) => {
              return `${xx},${yy}`
            })
            .filter(k => {
              return !edgeNodes.has(k)
            })
          edgeNodes.set(nodes[i], {
            distance: nextDist,
            nodes: nextNodes
          })
        } else {
          remainingNodes.push(nodes[i])
        }
      }
      if (remainingNodes.length > 0) {
        edgeNodes.set(node, { distance, nodes: remainingNodes})
      } else {
        edgeNodes.delete(node)
      }
    }
  }
}

// Part 1:
// console.log(findShortestPath(input))

// Part 2:
const graph = tileGraph(input)
console.log(findShortestPath(graph))
