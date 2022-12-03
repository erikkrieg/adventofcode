from std/strutils import parseInt
from std/algorithm import sort
from std/math import sum

proc day01() =
  echo("- Day 01")
  let input = open("input/day-01.txt")
  defer: input.close()

  var sums = @[0]
  var line : string
  while input.readLine(line):
    case line:
      of "":
        sums.add(0)
      else:
        sums[^1] += parseInt(line)
  sums.sort()
  sums = sums[^3..^1]

  echo("  - Part 1: Most calories is ", sums.max())
  echo("  - Part 2: Sum of calories for top 3 elves is ", sums.sum())

when isMainModule:
  echo("Advent of Code 2022")
  day01()

