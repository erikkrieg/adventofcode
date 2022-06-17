use std::fs;

fn main() {
    let inputs = fs::read_to_string("./inputs").expect("./inputs file not found");
    let (x, y) = part_1(&inputs);
    println!("Part 1: {:?}", x * y);
    let (x, y) = part_2(&inputs);
    println!("Part 2: {:?}", x * y);
}

fn part_1(commands: &str) -> (i32, i32) {
    let mut x = 0;
    let mut y = 0;

    commands.lines().for_each(|command| {
        let mut cmd_parts = command.split(' ');
        let command = cmd_parts.next().unwrap();
        let value = cmd_parts.next().unwrap().parse::<i32>().unwrap();
        match command {
            "forward" => x += value,
            "down" => y += value,
            "up" => y -= value,
            _ => panic!("command not found: {}", command),
        }
    });

    (x, y)
}

// Rust fmt is remove a comma from the multiline
// match statement :(
#[rustfmt::skip]
fn part_2(commands: &str) -> (i32, i32) {
    let mut aim = 0;
    let mut x = 0;
    let mut y = 0;

    commands.lines().for_each(|command| {
        let mut parts = command.split(' ');
        let command = parts.next().unwrap();
        let value = parts.next().unwrap().parse::<i32>().unwrap();
        match command {
            "forward" => {
                x += value;
                y += value * aim;
            },
            "down" => aim += value,
            "up" => aim -= value,
            _ => panic!("command not found: {}", &command),
        }
    });

    (x, y)
}
