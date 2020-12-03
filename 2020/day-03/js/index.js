const fs = require('fs')
const countTreesOnRoute = require('./countTreesOnRoute')

const mapMatrix = fs.readFileSync('../data', 'utf8').split('\n')
const routes = [[1, 1], [3, 1], [5, 1], [7, 1], [1, 2]]

console.log('Count of trees for each route:')
const productOfTrees = routes.reduce((product, route) => {
  const trees = countTreesOnRoute(mapMatrix, route)
  console.log(` - [${route}]:`, trees)
  return product * trees
}, 1)

console.log('Product of trees from each route:', productOfTrees)
