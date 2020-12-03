module.exports = (policy, password) => {
  const [positions, char] = policy.split(' ')
  const chars = new Set(positions.split('-').map(pos => password[parseInt(pos, 10) - 1]))
  return chars.has(char) && chars.size === 2
}
