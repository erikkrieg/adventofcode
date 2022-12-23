let part_one = stream => {
  let largest_sum = ref(0);
  let current_sum = ref(0);
  stream
  |> Stream.iter(line => {
       switch (line) {
       | "" =>
         largest_sum := max(largest_sum^, current_sum^);
         current_sum := 0;
       | _ => current_sum := current_sum^ + int_of_string(line)
       }
     });
  Printf.printf("  - Part 1: %d\n", largest_sum^);
};

let solve = () => {
  print_endline("- Day 01:");
  let input = open_in("input/day-01.txt");
  let stream =
    Stream.from(_ => {
      switch (input_line(input)) {
      | line => Some(line)
      | exception End_of_file => None
      }
    });
  part_one(stream);
};
