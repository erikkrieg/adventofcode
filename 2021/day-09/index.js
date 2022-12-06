const fs = require("fs")
const { setgroups } = require("process")
const stdinBuffer = fs.readFileSync(0)
const inputs = stdinBuffer.toString().split('\n')
// const inputs = `2199943210
// 3987894921
// 9856789892
// 8767896789
// 9899965678`.split('\n')

// let riskLevel = 0
// const lowPoints = []
// for (let r = 0; r < inputs.length; r++) {
//     const row = inputs[r]
//     for (let c = 0; c < row.length; c++) {
//         const smallerOrEqAdj = [[r-1, c], [r+1, c], [r, c-1], [r, c + 1]].reduce((count, [ri, ci]) => {
//             const height = inputs?.[ri]?.[ci] || 9
//             return count + (row[c] >= height)
//         }, 0)
//         if (smallerOrEqAdj === 0) {
//             riskLevel += (row[c] - -1)
//             lowPoints.push([r,c])
//         }
//     }
// }
// console.log({riskLevel})

const heightMap = inputs.map((hm) => {
    return hm.split('').map(h => ({ height: h, basin: null }))
    
})
let riskLevel = 0
const lowPoints = []
const basinsSet = new Set()
for (let r = 0; r < heightMap.length; r++) {
    const row = heightMap[r]
    for (let c = 0; c < row.length; c++) {
        const smallerOrEqAdj = [[r-1, c], [r+1, c], [r, c-1], [r, c + 1]].reduce((count, [ri, ci]) => {
            const adj = heightMap[ri]?.[ci]
            if (!adj || adj.height == '9') return count
            if (row[c].height !== '9') {
                // const basin = row[c].basin || new Set([`${r},${c}`])
                // row[c].basin = basin
                // adj.basin = basin
                // basin.add(`${ri},${ci}`)
                // const basin = adj.basin || row[c].basin || new Map([[`${r},${c}`, row[c].height]])
                // row[c].basin = basin
                // adj.basin = basin
                // basin.set(`${ri},${ci}`, adj.height)

                // const basin = new Map([
                //     ...(row[c].basin || [[`${r},${c}`, row[c].height]]),
                //     ...(adj.basin || [[`${ri},${ci}`, adj.height]]),
                // ])
                // row[c].basin = basin
                // adj.basin = basin

                if (adj.basin && row[c].basin) {
                    // pass basin from larger height
                    // to smaller?
                }
                
            }

            // if (adj.height < row[c].height) {
            //     adj.basinVal++
            //     return count + 1
            // } else if (row[c].height === adj.height){
            //     return count +1
            // }
            // return count
            const adjHeight = heightMap[ri]?.[ci]?.height || 9
            return count + (row[c].height >= adjHeight)
        }, 0)
        if (smallerOrEqAdj === 0) {
            riskLevel += (row[c].height - -1)
            lowPoints.push([r,c])
        }
    }
}
console.log({riskLevel})

// lowPoints
//     .map(([r,c]) => new Map([...heightMap[r][c].basin].sort((a,b) => a[0].localeCompare(b[0]))))
//     .forEach(basin => console.log(basin.size))
// console.log(basinsSet)


// const [r,i] = lowPoints[0]
// console.log(heightMap[r][i])

// heightMap.forEach(hm => hm.forEach(h => console.log(h.basin)))

const sortSet = (s) => new Set([...s].sort((a,b) => a.localeCompare(b)))
const adj = (([r,c]) => [[r-1, c], [r,c], [r - -1, c], [r, c-1], [r, c - -1]].map((a) => a.toString()))
const basins = []
for (let i = 0; i < lowPoints.length; i++) {
    const [r, c] = lowPoints[i]
    const notChecked = new Set(adj([r,c]))
    const checked = new Set()
    const basin = new Set()
    const walls = new Set()

    while (notChecked.size > 0) {
        const values = notChecked.values() 
        // console.log(values)
        for (const nc of values) {
            
            const [rr, cc] = nc.split(',')
            const height = heightMap[rr]?.[cc]?.height || 9
            if (height < 9) {
                // console.log(nc, height, adj([rr,cc]).filter((a) => !checked.has(a)))
                basin.add(nc)
                adj([rr,cc])
                    .filter((a) => !checked.has(a))
                    .forEach((a) => notChecked.add(a))
            } else {
                walls.add(nc)
            }
            checked.add(nc)
            notChecked.delete(nc)
        }
    }
    // console.log(sortSet(basin))
    basins.push(basin)
}

const productOfLargestBasins = basins.map(b => b.size).sort((a,b) => a-b).slice(-3).reduce((product, factor) => product * factor, 1)
console.log({productOfLargestBasins})

// // this is approach is going outward from basins and finding all the paths until a 9 is hit
// const basinSize = ([r,c]) => {
//     // const smallerOrEqAdj = [[r-1, c], [r+1, c], [r, c-1], [r, c + 1]].reduce((count, [ri, ci]) => {
//     //     const height = inputs?.[ri]?.[ci] || 9
//     //     return count + (row[c] >= height)
//     // }, 0)

//     const basin = new Set([[r,c].toString()])
//     console.log(basin)
//     return 2
// }
// const productOfLargestBasins = lowPoints.map(basinSize).sort((a,b) => a-b).slice(-3).reduce((product, factor) => product * factor, 1)
// console.log({productOfLargestBasins})
