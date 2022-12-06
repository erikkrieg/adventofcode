const fs = require("fs")
const { arrayBuffer } = require("stream/consumers")
const stdinBuffer = fs.readFileSync(0)
const input = stdinBuffer.toString()

const extractInput = (i) => {
  const [imgEnhanceAlgo, img] = i.split('\n\n')
  return {
    imgEnhanceAlgo,
    img: img.split('\n')
  }
}

const padImg = (img, pad=3, char='.') => {
  const xLen = img[0].length + (pad * 2)
  const emptyRow = (new Array(xLen).fill(char)).join('')
  for (let i = 0; i < img.length; i++) {
    img[i] = `${emptyRow.slice(0, pad)}${img[i]}${emptyRow.slice(0, pad)}`
  }
  for (let i = 0; i < pad; i++) {
    img.unshift(emptyRow)
    img.push(emptyRow)
  }
  return img
}



const enhance = ({ imgEnhanceAlgo, img }, times=1) => {
  let char = '.'
  if (imgEnhanceAlgo[0] === '#' && times % 2 === 1) {
    char = '#'
  }
  img = padImg(img, 1, char)
  const newImg = img.slice(0).map(i => i.split(''))
  for (let i = 0; i < img.length; i++) {
    for (let r = 0; r < img[i].length; r++) {
      const bin = [
        img[i-1]?.[r-1],  img[i-1]?.[r],  img[i-1]?.[r+1],
        img[i]?.[r-1],    img[i]?.[r],    img[i]?.[r+1],
        img[i+1]?.[r-1],  img[i+1]?.[r],  img[i+1]?.[r+1],
      ].map(b => {
        if (b === '#') return 1
        else if (b === '.') return 0
        return char === '#' ? 1 : 0
      }).join('')
      const int = parseInt(bin, 2)
      newImg[i][r] = imgEnhanceAlgo[int]
    }
  }
  const outputImg = newImg.map(i => i.join(''))
  return --times > 0 ? enhance({ imgEnhanceAlgo, img: outputImg }, times) : outputImg
}

// const { imgEnhanceAlgo, img } = extractInput(input)
// const x2 = enhance({ imgEnhanceAlgo, img: padImg(img, 100) }, 50)
const x2 = enhance(extractInput(input), 2)
const litPixels = x2.reduce((total, row) => {
  return total + (row.match(/#/g)?.length || 0)
}, 0)

;[
  // [extractInput(input)],
  // [JSON.stringify(padImg(extractInput(input).img), null, 2)],
  [JSON.stringify(x2, null, 2)],
  [litPixels],
].forEach((out) => console.log(...out))
