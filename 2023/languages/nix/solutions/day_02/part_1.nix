with builtins;
let
  # Imports
  inherit (import <nixpkgs/lib>) toInt splitString;
  inherit (import ../../lib) debug lines sum;

  # Test data
  test = (getEnv "TEST") == "true";
  test_input = ''
    Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
    Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
    Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
    Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
    Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green
  '';

  count_balls = r: color:
    let
      round = debug r;
      found = match ".* ([0-9]+) ${color}.*" round;
      count = if debug found == null then 0 else toInt (head found);
    in
    count;

  valid_round = r:
    let
      red = count_balls r "red" <= 12;
      green = count_balls r "green" <= 13;
      blue = count_balls r "blue" <= 14;
    in
    red && green && blue;

  parse_game = g:
    let
      rounds = debug splitString ";" g;
      valid_rounds = map valid_round rounds;
      game_id = toInt (head (match "Game ([0-9]+).*" g));
    in
    if any (x: !x) valid_rounds then 0 else game_id;

  inputs = lines (if test then test_input else readFile ../../inputs/day-2.txt);
  possible_games = debug map parse_game inputs;
  result = sum possible_games;
in
result
