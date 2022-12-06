from std/sequtils import deduplicate

# The crux of this puzzle for me was figuring out how I could write procedures
# that accept arrays of different lengths. I was using arrays because it appeared
# to be the data type that should be used with `readChars`, which I wanted to
# try using.
# I was worried that openArray might not behave correctly with readChars, but
# it fortunately did respect the array's capacity. Unfortunately I made a bad
# assumption that _it would not_ and took a fruitless detour into macros XD

proc shiftMarker(arr: var openArray[char], insert: char) =
  for a in (arr.low + 1) .. arr.high():
    arr[a-1] = arr[a]
  arr[arr.high()] = insert

proc findMarker(marker: var openArray[char]): int =
  let input = open("input/day-06.txt")
  defer: input.close()
  var
    iterations = marker.len()
    cursor: array[1, char]
  discard readChars(input, marker)
  while readChars(input, cursor) > 0:
    if deduplicate(marker).len() == marker.len():
      break
    shiftMarker(marker, cursor[0])
    iterations += 1
  return iterations

proc solve*() =
  echo("- Day 06")
  var
    partOneArr: array[4, char]
    partTwoArr: array[14, char]
  echo("  - Part 1: ", findMarker(partOneArr))
  echo("  - Part 2: ", findMarker(partTwoArr))

