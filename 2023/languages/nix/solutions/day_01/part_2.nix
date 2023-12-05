with builtins;
let
  # Imports
  lib = import <nixpkgs/lib>;

  # Functions
  inherit (lib) toInt;
  toDigit = v:
    let
      table = {
        one = "1";
        two = "2";
        three = "3";
        four = "4";
        five = "5";
        six = "6";
        seven = "7";
        eight = "8";
        nine = "9";
      };
    in
    if hasAttr v table then table.${v} else v;
  debug = v: lib.traceIf ((getEnv "DEBUG") == "true") (deepSeq v v) v;
  lines = s: filter (x: (!isList x) && (x != "")) (split "\n" s);
  pattern = ".*(one|two|three|four|five|six|seven|eight|nine|[1-9])";
  firstDigitRec = str: end:
    let
      sub = substring 0 end str;
      match_digit = match pattern sub;
    in
    if isNull match_digit then firstDigitRec str (end + 1) else match_digit;
  firstDigit = s: firstDigitRec s 0;
  lastDigit = s: match "${pattern}.*" s;
  sum = ints: foldl' (a: b: a + b) 0 ints;

  # Test data
  test = (getEnv "TEST") == "true";
  test_input = ''
    two1nine
    eightwothree
    abcone2threexyz
    xtwone3four
    4nineeightseven2
    zoneight234
    7pqrstsixteen
  '';

  inputs = debug (lines (if test then test_input else readFile ../../inputs/day-1.txt));
  digits = debug (map (x: toInt (toDigit (head (firstDigit x)) + toDigit (head (lastDigit x)))) inputs);
  result = sum digits;
in
result
