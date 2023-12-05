with builtins;
let
  # Imports
  lib = import <nixpkgs/lib>;

  # Functions:
  inherit (lib.strings) toInt;
  debug = v: lib.traceIf ((getEnv "DEBUG") == "true") (deepSeq v v) v;
  lines = s: filter (x: (!isList x) && (x != "")) (split "\n" s);
  firstDigit = s: match "[a-z]*([1-9]).*" s;
  lastDigit = s: match ".*([1-9]).*" s;
  sum = ints: foldl' (a: b: a + b) 0 ints;

  # Test data
  test = (getEnv "TEST") == "true";
  test_input = ''
    1abc2
    pqr3stu8vwx
    a1b2c3d4e5f
    treb7uchet
  '';

  inputs = lines (if test then test_input else readFile ../../inputs/day-1.txt);
  digits = debug (map (x: toInt (head (firstDigit x) + head (lastDigit x))) inputs);
  result = sum digits;
in
result
