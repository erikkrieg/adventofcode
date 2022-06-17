const fs = require("fs")
const stdinBuffer = fs.readFileSync(0)
const inputs = stdinBuffer.toString().split('\n').map(i => i.split(''));

const makeChunk = (line, start) => {
    const tokenMap = { '(': '()', '[': '[]', '{': '{}', '<': '<>' }
    const chunk = {
        type: tokenMap[line[start]],
        start,
        end: undefined,
        chunks: []
    }
    
    let pointer = start
    while (!chunk.end) {
        pointer++
        const c = line[pointer]
        if (!c) {
            line.push(chunk.type[1])
            chunk.end = pointer
        } else if (tokenMap[c]) {
            const childChunk = makeChunk(line, pointer)
            pointer = childChunk.end
            chunk.chunks.push(childChunk)
        } else if (c === chunk.type[1]) {
            chunk.end = pointer
        } else {
            // this is an invalid token!
            throw ({ type: 'INVALID_TOKEN', data: c })
        }
    }
    return chunk
}
const getAutocompleteScores = (lines) => {
    const scoreMap = { ')': 1, ']': 2, '}': 3, '>': 4 }
    const scores = []
    for (let l = 0; l < lines.length; l++) {
        const initialLength = lines[l].length
        let pointer = 0
        while (pointer < lines[l].length) {
            let chunk;
            try {
                chunk = makeChunk(lines[l], pointer)
            } catch (err) {
                if (err.type === 'INVALID_TOKEN') {
                    break;
                }
            }
            if (!chunk) break
            pointer = chunk.end + 1
        }
        if (lines[l].length !== initialLength) {
            // console.log('added tokens:',lines[l].slice(initialLength).join(''))
            const autoTokens = lines[l].slice(initialLength)
            scores.push(autoTokens.reduce((s, t) => s * 5 + scoreMap[t], 0))
        }
    }
    return scores
}

const getMedianScore = (scores) => {
    scores.sort((a,b) => a-b)
    return scores[scores.length / 2 | 0]
}

const test = `[({(<(())[]>[[{[]{<()<>>
[(()[<>])]({[<{<<[]>>(
{([(<{}[<>[]}>{[]{[(<()>
(((({<>}<{<{<>}{[]{[]{}
[[<[([]))<([[{}[[()]]]
[{[{({}]{}}([{[{{{}}([]
{<[[]]>}<{[{[{[]{()[[[]
[<(<(<(<{}))><([]([]()
<{([([[(<>()){}]>(<<{{
<{([{{}}[<[[[<>{}]]]>[]]`.split('\n').map(i => i.split(''))
console.log('Test score:', getMedianScore(getAutocompleteScores(test)))
console.log('Test score:', getMedianScore(getAutocompleteScores(inputs)))
