module.exports = (map, [dx, dy]) => {
  const mapSize = { x: map[0].length, y: map.length }
  const pos = { x: 0, y: 0 }
  let treeCount = 0
  while (pos.y < mapSize.y - 1) {
    pos.x = (pos.x + dx) % mapSize.x
    pos.y += dy
    treeCount += map[pos.y][pos.x] === '#'
  }
  return treeCount
}
