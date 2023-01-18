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

fn lines(reader: BufReader<File>) -> impl Iterator<Item = String> {
    reader.lines().into_iter().filter_map(|line| line.ok())
}

fn part_one(reader: BufReader<File>) {
    let priority_sum = lines(reader).fold(0, |acc, x| {
        let mid = x.len() / 2;
        acc + score(
            &first_common_char(&[&x[0..mid], &x[mid..]]).expect("No common character was found"),
        )
    });
    println!("  - Part 1: {priority_sum}");
}

fn part_two(reader: BufReader<File>) {
    let mut priority_sum = 0;
    let iter = lines(reader);
    for arr in iter.array_chunks::<3>() {
        let group = arr.iter().map(|s| s.as_str()).collect::<Vec<&str>>();
        priority_sum += score(&first_common_char(&group).expect("No common character was found"))
    }
    println!("  - Part 2: {priority_sum}");
}

fn first_common_char(strs: &[&str]) -> Option<char> {
    match strs {
        [first] => first.chars().next(),
        [first, rest @ ..] => first
            .chars()
            .into_iter()
            .find(|&c| rest.iter().all(|r| r.contains(c))),
        _ => None,
    }
}

fn score(c: &char) -> u32 {
    let c = *c as u32;
    if c >= 97 {
        c - 96
    } else {
        c - 38
    }
}
