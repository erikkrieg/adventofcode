use std::collections::BinaryHeap;
use std::fs::File;
use std::io::{self, prelude::*, BufReader};

pub fn solve() -> io::Result<()> {
    println!("- Day 01");

    let sums = calorie_sums()?;

    // Solve part 1
    println!("  - Part 1: Most calories is {}", sums.peek().unwrap_or(&0));

    // Solve part 2
    println!(
        "  - Part 2: Sum of calories for top 3 elves is {}",
        sums.iter().take(3).sum::<u32>()
    );

    Ok(())
}

fn calorie_sums() -> Result<BinaryHeap<u32>, io::Error> {
    let file = File::open(&"input/day-01.txt")?;
    let reader = BufReader::new(file);
    Ok(reader
        .lines()
        .into_iter()
        .fold((0, BinaryHeap::new()), |mut acc, x| {
            let x = x.unwrap();
            if x.is_empty() {
                acc.1.push(acc.0);
                return (0, acc.1);
            }
            acc.0 += x.parse::<u32>().unwrap_or(0);
            acc
        })
        .1)
}
