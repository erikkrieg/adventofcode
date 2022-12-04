import std/sequtils

proc score(c: char): int =
    let c = int(c)
    if c >= 97:
      # a-z is scored 1-26
      return c - 96 
    else:
      # A-Z is scored 27-52
      return c - 38

proc partOne() =
  let input = open("input/day-03.txt")
  var
    line: string
    prioritySum: int
  while input.readLine(line):
    let split = toSeq(line.items).distribute(2)
    let match = filter(split[0], proc(c: char): bool = split[1].contains(c))
    prioritySum += score(match[0])
  echo("  - Part 1: ", prioritySum)

proc partTwo() =
  let input = open("input/day-03.txt")
  var
    line: string
    prioritySum: int
    group: seq[seq[char]]
  while input.readLine(line):
    group.add(deduplicate(toSeq(line.items)))
    if group.len() == 3:
      let match = filter(
        group[0],
        proc(c: char): bool = group[1].contains(c) and group[2].contains(c)
      )
      prioritySum += score(match[0])
      group = @[]
  echo("  - Part 1: ", prioritySum)


proc solve*() =
  echo("- Day 03")
  partOne()
  partTwo()

