with builtins;
let
  lib = import <nixpkgs/lib>;
in
{
  debug = v: lib.traceIf ((getEnv "DEBUG") == "true") (deepSeq v v) v;
  lines = s: filter (x: (!isList x) && (x != "")) (split "\n" s);
  sum = ints: foldl' (a: b: a + b) 0 ints;
}
