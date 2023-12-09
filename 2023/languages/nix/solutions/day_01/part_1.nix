with builtins;
let
  # Imports
  inherit (import <nixpkgs/lib>) toInt;
  inherit (import ../../lib) debug lines sum;

  # Test data
  test = (getEnv "TEST") == "true";
  test_input = ''
    1abc2
    pqr3stu8vwx
    a1b2c3d4e5f
    treb7uchet
  '';

  firstDigit = s: match "[a-z]*([1-9]).*" s;
  lastDigit = s: match ".*([1-9]).*" s;

  inputs = lines (if test then test_input else readFile ../../inputs/day-1.txt);
  digits = debug (map (x: toInt (head (firstDigit x) + head (lastDigit x))) inputs);
  result = sum digits;
in
result
