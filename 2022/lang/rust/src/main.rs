use std::fs::File;
use std::io::{self, prelude::*, BufReader};

fn main() -> io::Result<()> {
    println!("Advent of Code 2022");
    println!("- Day 01");

    let file = File::open(&"input/day-01.txt")?;
    let reader = BufReader::new(file);
    let mut calorie_sums = Vec::<u32>::new();
    let mut current_calorie_sum = 0;
    for line in reader.lines() {
        let line = line?;
        if line.is_empty() {
            calorie_sums.push(current_calorie_sum);
            current_calorie_sum = 0;
        } else {
            let num = line.parse::<u32>().unwrap();
            current_calorie_sum += num;
        }
    }
    calorie_sums.sort();

    // Solve part 1
    println!(
        "  - Part 1: Most calories is {}",
        calorie_sums.last().unwrap_or(&0)
    );

    // Solve part 2
    let top_sums: u32 = calorie_sums.iter().rev().take(3).sum();
    println!(
        "  - Part 2: Sum of calories for top 3 elves is {}",
        &top_sums
    );

    Ok(())
}
