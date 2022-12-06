const fs = require("fs")
const stdinBuffer = fs.readFileSync(0)
const inputs = stdinBuffer.toString().split('\n'); /*inputs.pop();*/

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
           throw ({ type: 'INCOMPLETE_CHUNK', data: chunk }) 
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
const getSyntaxErrorScore = (lines) => {
    const scoreMap = { ')': 3, ']': 57, '}': 1197, '>': 25137 }
    let score = 0
    for (let l = 0; l < lines.length; l++) {
        let pointer = 0
        while (pointer < lines[l].length) {
            let chunk;
            try {
                chunk = makeChunk(lines[l], pointer)
            } catch (err) {
                if (err.type === 'INVALID_TOKEN') {
                    // console.log(lines[l], err, scoreMap[err.data])
                    score += scoreMap[err.data]
                }
            }
            if (!chunk) break
            pointer = chunk.end + 1
        }
    }
    return score
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
<{([{{}}[<[[[<>{}]]]>[]]`.split('\n')
console.log('Test error score:', getSyntaxErrorScore(test))
console.log('Input error score:', getSyntaxErrorScore(inputs))
