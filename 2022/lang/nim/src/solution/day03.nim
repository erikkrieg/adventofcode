import std/sequtils

proc solve*() =
  echo("- Day 03")
  let input = open("input/day-03.txt")

  var
    line: string
    prioritySum: int

  while input.readLine(line):
    let split = toSeq(line.items).distribute(2)
    let match = filter(split[0], proc(s: char): bool = split[1].contains(s))
    let c = int(match[0])
    if c >= 97:
      # a-z is scored 1-26
      prioritySum += c - 96 
    else:
      # A-Z is scored 27-52
      prioritySum += c - 38

  echo("  - Part 1: ", prioritySum)
