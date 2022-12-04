use std::collections::HashSet;
use std::fs::File;
use std::io::{self, prelude::*, BufReader};

pub fn solve() -> io::Result<()> {
    println!("- Day 03");

    let file = File::open("input/day-03.txt")?;
    part_one(BufReader::new(file));

    let file = File::open("input/day-03.txt")?;
    part_two(BufReader::new(file));

    Ok(())
}

fn part_one(reader: BufReader<File>) {
    let priority_sum = reader
        .lines()
        .into_iter()
        .filter_map(|r| r.ok())
        .fold(0, |mut acc, x| {
            let a: HashSet<char> = x[0..x.len() / 2].chars().collect();
            let b: HashSet<char> = x[x.len() / 2..].chars().collect();
            for i in a.intersection(&b) {
                acc += score(i);
            }
            acc
        });
    println!("  - Part 1: {priority_sum}");
}

fn part_two(reader: BufReader<File>) {
    let mut priority_sum = 0;
    let iter = reader.lines().into_iter().filter_map(|r| r.ok());
    for arr in iter.array_chunks::<3>() {
        let sets: Vec<HashSet<char>> = arr.iter().map(|x| x.chars().collect()).collect();
        for i in sets[0].intersection(&sets[1]) {
            if sets[2].contains(i) {
                priority_sum += score(i)
            }
        }
    }
    println!("  - Part 2: {priority_sum}");
}

fn score(c: &char) -> u32 {
    let c = *c as u32;
    if c >= 97 {
        c - 96
    } else {
        c - 38
    }
}
