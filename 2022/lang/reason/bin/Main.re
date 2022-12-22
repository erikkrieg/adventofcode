print_endline("Advent of Code 2022");
switch (Sys.argv) {
  | [| _, day |] when day == "1" => Lib.Day01.solve()
  | [| _, day |] => print_endline("No solution found for the specified day: " ++ day)
  | _ => print_endline("Please provide the date for the puzzle to solve: i.e. `esy run 1`")
};

