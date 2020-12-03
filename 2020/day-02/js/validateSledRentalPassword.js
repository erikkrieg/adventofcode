module.exports = (policy, password) => {
  const [range, char] = policy.split(' ')
  const [min, max] = range.split('-').map(s => parseInt(s, 10))

  let matchedCharCount = 0
  for (let index = 0; index < password.length; index++) {
    if (password[index] === char) matchedCharCount++
  }

  return matchedCharCount >= min && matchedCharCount <= max
}
