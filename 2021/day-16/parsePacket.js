const fs = require("fs")
const stdinBuffer = fs.readFileSync(0)
const input = stdinBuffer.toString()

const hexToBin = (hex) => {
  const conversionTable = {
    "0": "0000",
    "1": "0001",
    "2": "0010",
    "3": "0011",
    "4": "0100",
    "5": "0101",
    "6": "0110",
    "7": "0111",
    "8": "1000",
    "9": "1001",
    "A": "1010",
    "B": "1011",
    "C": "1100",
    "D": "1101",
    "E": "1110",
    "F": "1111",
  }
  
  let binary = ''
  for (let i = 0; i < hex.length; i++) {
    binary += conversionTable[hex[i]]
  }
  return binary 
}

// {
//   headers,
//   value: [
//     { headers, value: 5 }
//   ]
// }
const cut = (arr, n) => {
  return {
    rest: arr.slice(n),
    cut: arr.slice(0,n)
  }
}
const parse = (transmission) => {
  let bin = transmission
  while (bin.length > 0) {
    
  }
  let { rest, cut: version } = cut(transmission, 3)
  let { rest, cut: type } = cut(rest, 3)
}


[
  parse(hexToBin('D2FE28')),
].forEach((out) => {
  console.log(...out)
})
