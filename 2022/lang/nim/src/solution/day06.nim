from std/sequtils import deduplicate

proc shiftMarker(arr: array[4, char], insert: char): array[4, char] =
  for a in 1..3:
    result[a-1] = arr[a]
  result[3] = insert

proc endOfMarkerIndex(): int =
  let input = open("input/day-06.txt")
  defer: input.close()
  var
    iterations = 4
    potentialMarker: array[4, char]
    cursor: array[1, char]
  discard readChars(input, potentialMarker)
  while readChars(input, cursor) > 0:
    if deduplicate(potentialMarker).len() == 4:
      break
    potentialMarker = shiftMarker(potentialMarker, cursor[0])
    iterations += 1
  return iterations

proc partOne() =
  echo("  - Part 1: ", endOfMarkerIndex())

proc solve*() =
  echo("- Day 06")
  partOne()

