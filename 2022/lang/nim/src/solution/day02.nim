import std/tables

# I wanted to try solving with this with different language features to get more
# exposure to the language, so I used hashmap and control flow (case).

const decoder: Table[string, uint] = {
  "A X": uint(3 + 0),
  "B X": uint(1 + 0),
  "C X": uint(2 + 0),

  "A Y": uint(1 + 3),
  "B Y": uint(2 + 3),
  "C Y": uint(3 + 3),

  "A Z": uint(2 + 6),
  "B Z": uint(3 + 6),
  "C Z": uint(1 + 6),
  }.toTable()

proc part1Score(code: string): uint =
  case code:
     of "A X":
       result = 1 + 3
     of "B X":
       result = 1 + 0
     of "C X":
       result = 1 + 6
     of "A Y":
       result = 2 + 6
     of "B Y":
       result = 2 + 3
     of "C Y":
       result = 2 + 0
     of "A Z":
       result = 3 + 0
     of "B Z":
       result = 3 + 6
     of "C Z":
       result = 3 + 3

proc part2Score(code: string): uint =
  return decoder[code]

proc solve*() =
  echo("- Day 02")
  let input = open("input/day-02.txt")
  defer: input.close()

  var
    totalPart1Score: uint = 0
    totalPart2Score: uint = 0
    line: string

  while input.readLine(line):
    totalPart1Score += part1Score(line) 
    totalPart2Score += part2Score(line) 
       
  echo("  - Part 1:", totalPart1Score)
  echo("  - Part 2:", totalPart2Score)

