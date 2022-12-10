use std::{
    fs::File,
    io::{self, BufRead, BufReader},
};

pub fn solve() -> io::Result<()> {
    println!("- Day 10:");
    let input = File::open("input/day-10.txt")?;
    let reader = BufReader::new(input);
    let cycles = calculate_cycles(reader);
    part_one(&cycles);
    part_two(&cycles);
    Ok(())
}

fn calculate_cycles(commands: BufReader<File>) -> Vec<i32> {
    let (cycles, _): (Vec<i32>, i32) =
        commands
            .lines()
            .filter_map(|l| l.ok())
            .fold((vec![1], 0), |(mut cycles, add), l| {
                cycles.push(*cycles.last().unwrap() + add);
                let add = match l.split_whitespace().nth(1) {
                    Some(addend) => {
                        // addx uses an extra cycle
                        cycles.push(*cycles.last().unwrap());
                        addend.parse().expect("addx command only accepts integers")
                    }
                    None => 0,
                };
                (cycles, add)
            });
    cycles
}

fn part_one(cycles: &[i32]) {
    let signal_strength_sum: i32 = cycles
        .iter()
        .enumerate()
        .skip(20)
        .step_by(40)
        .map(|(i, v)| i as i32 * v)
        .sum();
    println!("  - Part 1: {signal_strength_sum}");
}

fn part_two(cycles: &[i32]) {
    let screen: Vec<String> = cycles
        .iter()
        .skip(1)
        .enumerate()
        .map(|(i, c)| match (c - 1..=c + 1).contains(&(i as i32 % 40)) {
            true => "#".to_string(),
            false => ".".to_string(),
        })
        .collect();
    println!("  - Part 2:");
    screen
        .chunks(40)
        .for_each(|row| println!("{}", &row.join("")));
}
