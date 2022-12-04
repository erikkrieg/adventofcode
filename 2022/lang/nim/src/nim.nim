import os
import solution/day01
import solution/day02

when isMainModule:
  echo("Advent of Code 2022")
  let day = paramStr(1)
  case day
    of "01", "1": 
      day01.solve()
    of "02", "2": 
      day02.solve()
    else:
      echo("Solution not found: ", day)

