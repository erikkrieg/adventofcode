with builtins;
let
  test = true;
  lines = s: filter (x: (!isList x) && (x != "")) (split "\n" s);
  test_input = ''
    1abc2
    pqr3stu8vwx
    a1b2c3d4e5f
    treb7uchet
  '';
  inputs = lines (if test then test_input else readFile ./inputs/day-1.txt);
  result = builtins.map (x: (match "(1)" x)) inputs;
in
trace (head result) null
