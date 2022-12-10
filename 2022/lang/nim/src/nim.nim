import os
import solution/day01
import solution/day02
import solution/day03
import solution/day04
import solution/day06

when isMainModule:
  echo("Advent of Code 2022")
  let day = paramStr(1)
  case day
    of "01", "1": 
      day01.solve()
    of "02", "2": 
      day02.solve()
    of "03", "3": 
      day03.solve()
    of "04", "4": 
      day04.solve()
    of "06", "6": 
      day06.solve()
    else:
      echo("Solution not found: ", day)

