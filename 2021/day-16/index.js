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

const isValidPacket = ({ type, payload, lengthType, length}) => {
  return (type === '100' && payload) || (lengthType && length)
}

const parsePacket = (transmission) => {
  const decodedPackets = []
  let index = 0
  while (index < transmission.length) {
      const version = transmission.slice(index, index+=3)
      const type = transmission.slice(index, index+=3)
      const payloadStart = index
      let payload, lengthType, length;
      if (type === '100') {
        while (transmission[index] === '1') {
          index += 5
        }
        payload = transmission.slice(payloadStart, index+=5)
      } else {
        lengthType = transmission[index++]
        if (lengthType == 0) {
          length = transmission.slice(index, index += 15)
        } else if (lengthType == 1) {
          length = transmission.slice(index, index += 11)
        }
      }
      const packet = { version, type, lengthType, length, payload }
      if (isValidPacket(packet)) {
        decodedPackets.push(packet)
      }
  }
  return decodedPackets
}

const reducer = (sum, { packet, subpackets }) => {
  for (let key in packet) {
    sum += packet[key]?.length || 0
  }
  return subpackets ? subpackets.reduce(reducer, sum) : sum
}

const nestPackets = (packets, pointer=0) => {
  // console.log('start', {pointer})
  const subpackets = []
  const packet = packets[pointer]
  if (!packet) {
    // console.log('no packet', {pointer})
    return
  }
  if (packet.type === '100') {
    // console.log('return literal')
    return { packet, nextPointer: pointer + 1 }
  } else {
    let remainingSubpackets = true
    pointer++
    while (remainingSubpackets) {
      // console.log('LF sub')
      const res = nestPackets(packets, pointer)
      if (!res) {
        // console.log('no res')
        remainingSubpackets = false
        break
      }
      // console.log({res})
      subpackets.push(res)
      pointer = res.nextPointer
      if (packet.lengthType === '0') {
        const subpacketLen = subpackets.reduce(reducer, 0) 
        remainingSubpackets = parseInt(packet.length, 2) > subpacketLen
        // console.log( packet.version + '-' + packet.type, 'remaining:', remainingSubpackets, subpacketLen)
      } else if (packet.lengthType === '1') {
        remainingSubpackets = parseInt(packet.length, 2) > subpackets.length
      }
    }
  }
  // console.log(packet.version + '-' + packet.type, 'returning')
  return { packet, subpackets, nextPointer: pointer }
}


// console.log(JSON.stringify(parsePacket2(hexToBin('D2FE28')), 0, 2))
// console.log(JSON.stringify(parsePacket2(hexToBin('38006F45291200')), 0, 2))
// console.log(JSON.stringify(parsePacket(hexToBin('EE00D40C823060')), 0, 2))

// console.log(parsePacket(hexToBin('8A004A801A8002F478')).reduce((sum, { version }) => sum + parseInt(version, 2), 0))
// console.log(parsePacket(hexToBin('620080001611562C8802118E34')).reduce((sum, { version }) => sum + parseInt(version, 2), 0))
// console.log(parsePacket(hexToBin('620080001611562C8802118E34')).reduce((sum, { version }) => sum + parseInt(version, 2), 0))
// console.log(parsePacket(hexToBin('A0016C880162017C3686B18A3D4780')).reduce((sum, { version }) => sum + parseInt(version, 2), 0))
// console.log(parsePacket(hexToBin(input)).reduce((sum, { version }) => sum + parseInt(version, 2), 0))

// [
//   // nestPackets(parsePacket(hexToBin('38006F45291200'))),
//   // nestPackets(parsePacket(hexToBin('EE00D40C823060'))),
//   nestPackets(parsePacket(hexToBin('8A004A801A8002F478'))),
// ].forEach((out) => {
//   console.log(JSON.stringify(out, null, 2))
// })

const extractLiteralValue = (litPayload) => {
  // return litPayload.match(/.{5}/g).reduce((sum, v) => parseInt(v.slice(1), 2), 0)
  const literalValBin = litPayload.match(/.{5}/g).reduce((res, v) => {
    return res + v.slice(1)
  }, '')
  return parseInt(literalValBin, 2)
}
const evalPackets = ({ packet, subpackets }) => {
  const ops = {
    0: function sum (a, b) {
      a = isNaN(a) ? 0 : a
      return a + b
    },
    1: function product (a, b) {
      a = isNaN(a) ? 1 : a
      return a * b
    },
    2: function min (a, b) {
      a = isNaN(a) ? Infinity : a
      return Math.min(a, b)
    },
    3: function max (a, b) {
      a = isNaN(a) ? -Infinity : a
      return Math.max(a, b)
    },
    5: function greaterThan (a, b) {
      if (isNaN(a)) return b
      return a > b ? 1 : 0
    },
    6: function lessThan (a, b) {
      if (isNaN(a)) return b
      return a < b ? 1 : 0
    },
    7: function equalTo (a, b) {
      if (isNaN(a)) return b
      return a === b ? 1 : 0
    },
  }
  const type = parseInt(packet.type, 2)
  const op = ops[type]
  if (!subpackets) {
    console.log('no subp', packet)
  }
  return subpackets.reduce((val, sp) => {
    if (sp.packet.type === '100') {
      const payload = extractLiteralValue(sp?.packet?.payload)
      return op(val, payload)
    } else {
      return op(val, evalPackets(sp))
    }
  }, NaN)
}
const evalPacketsVersion = ({ packet, subpackets }) => {
  const version = parseInt(packet.version, 2)
  // console.log(packet, version)
  const type = parseInt(packet.type, 2)
  return subpackets.reduce((val, sp) => {
    const spVersion = parseInt(sp.packet.version, 2)
    if (sp.packet.type === '100') {
      // console.log(sp.packet, spVersion)
      return val + spVersion
    } else {
      return val + evalPacketsVersion(sp)
    }
  }, version)
}


[
  [evalPackets(nestPackets(parsePacket(hexToBin('C200B40A82')))), 'Expects:', 3],
  [evalPackets(nestPackets(parsePacket(hexToBin('04005AC33890')))), 'Expects:', 54],
  [evalPackets(nestPackets(parsePacket(hexToBin('880086C3E88112')))), 'Expects:', 7],
  [evalPackets(nestPackets(parsePacket(hexToBin('CE00C43D881120')))), 'Expects:', 9],
  [evalPackets(nestPackets(parsePacket(hexToBin('D8005AC2A8F0')))), 'Expects:', 1],
  [evalPackets(nestPackets(parsePacket(hexToBin('F600BC2D8F')))), 'Expects:', 0],
  [evalPackets(nestPackets(parsePacket(hexToBin('9C005AC2F8F0')))), 'Expects:', 0],
  [evalPackets(nestPackets(parsePacket(hexToBin('9C0141080250320F1802104A08')))), 'Expects:', 1],
  [evalPackets(nestPackets(parsePacket(hexToBin(input)))), 'Expects: ?'],
  

  // [JSON.stringify(nestPackets(parsePacket(hexToBin(input))), null, 2)],
  // [hexToBin(input)],
  // [hexToBin('32F5DF3B128')],
  // [evalPackets(nestPackets(parsePacket(hexToBin('32F5DF3B128')))), 'Expects: ?'],
  


  // [evalPacketsVersion(nestPackets(parsePacket(hexToBin('A0016C880162017C3686B18A3D4780'))))],
  // [evalPackets(nestPackets(parsePacket(hexToBin('A0016C880162017C3686B18A3D4780'))))],
  // [JSON.stringify(nestPackets(parsePacket(hexToBin('A0016C880162017C3686B18A3D4780'))), null, 2)],
  // [evalPacketsVersion(nestPackets(parsePacket(hexToBin('C0015000016115A2E0802F182340'))))],
  // [evalPacketsVersion(nestPackets(parsePacket(hexToBin('A0016C880162017C3686B18A3D4780'))))],
  // [evalPacketsVersion(nestPackets(parsePacket(hexToBin(input))))],
  // [evalPackets({
  //   "packet": {
  //     "version": "010",
  //     "type": "011",
  //     "lengthType": "1",
  //     "length": "00000000010"
  //   },
  //   "subpackets": [
  //     {
  //       "packet": {
  //         "version": "010",
  //         "type": "100",
  //         "payload": "01001"
  //       },
  //       "nextPointer": 3
  //     },{
  //       "packet": {
  //         "version": "010",
  //         "type": "000",
  //         "lengthType": "1",
  //         "length": "00000000011"
  //       },
  //       "subpackets": [
  //         {
  //           "packet": {
  //             "version": "010",
  //             "type": "100",
  //             "payload": "01001"
  //           },
  //           "nextPointer": 3
  //         },
  //         {
  //           "packet": {
  //             "version": "010",
  //             "type": "100",
  //             "payload": "01001"
  //           },
  //           "nextPointer": 3
  //         },
  //         {
  //           "packet": {
  //             "version": "010",
  //             "type": "100",
  //             "payload": "01001"
  //           },
  //           "nextPointer": 3
  //         }
  //       ]
  //     }
  //   ]
  // })],
].forEach((out) => {
  // console.log(JSON.stringify(out, null, 2))
  console.log(...out)
})