const fs = require("fs")
const stdinBuffer = fs.readFileSync(0)
const inputs = stdinBuffer.toString().split('\n')

// for each options recursively traverse. when at the end of path return it
// filter out paths that don't end in "end"

const log = (tar, out) => tar && console.log(out)

const findPaths = (edges, path, paths=[]) => {
  if (!path) {
    path = {
      uniq: new Set(['start']),
      nodes: ['start'],
      canVisitUniqNode: true
    }
  }
  const currentNode = path.nodes[path.nodes.length-1]
  if (currentNode === 'end') {
    paths.push(path)
    return paths
  }
  // could improve performance by caching this
  const nextNodes = edges.reduce((nodes, edge) => {
    if (edge.includes(currentNode)) {
      const nextNode = edge.split('-').filter(n => n !== currentNode)[0]
      if (nextNode !== 'start' && path.canVisitUniqNode || !path.uniq.has(nextNode)) {
        nodes.push(nextNode)
      }
    }
    return nodes
  }, [])
  if (nextNodes.length > 0) {
    nextNodes.forEach((nextNode) => {
      const newPath = {
        canVisitUniqNode: path.canVisitUniqNode,
        uniq:  new Set([...path.uniq]),
        nodes: [...path.nodes, nextNode]
      }
      if (nextNode != nextNode.toUpperCase()) {
        if (newPath.canVisitUniqNode && newPath.uniq.has(nextNode)) {
          newPath.canVisitUniqNode = false
        }
        newPath.uniq.add(nextNode)
      }
      findPaths(edges, newPath, paths)
    })
  } else {
    // This is a dead-end path
    // could include this in paths, or make an array for dead ends.
    // paths.push(path)
  }
  return paths
}

const paths = findPaths(inputs)
console.log(paths.length)
// console.log(paths.map(ps => JSON.stringify(ps.nodes)))