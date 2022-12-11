use core::ops::RangeInclusive;
use std::{
    fs::File,
    io::{self, BufRead, BufReader},
};

pub fn solve() -> io::Result<()> {
    println!("- Day 04:");
    let input = File::open("input/day-04.txt")?;
    let (full_overlap, partial_overlap) = calculate_overlap(BufReader::new(input));
    part_one(&full_overlap);
    part_two(&full_overlap, &partial_overlap);
    Ok(())
}

fn part_one(full_overlap: &Vec<Overlap>) {
    println!("  - Part 1: {}", full_overlap.len());
}

fn part_two(full_overlap: &Vec<Overlap>, partial_overlap: &Vec<Overlap>) {
    println!("  - Part 2: {}", full_overlap.len() + partial_overlap.len());
}

enum Overlap {
    Full,
    Partial,
    None,
}

trait Overlaps {
    fn overlaps(&self, other: &Self) -> Overlap;
}

impl Overlaps for RangeInclusive<i32> {
    fn overlaps(&self, other: &Self) -> Overlap {
        let overlap_size = self.end().min(other.end()) - self.start().max(other.start());
        let min_range_size = (self.end() - self.start()).min(other.end() - other.start());
        if overlap_size == min_range_size {
            Overlap::Full
        } else if overlap_size >= 0 {
            Overlap::Partial
        } else {
            Overlap::None
        }
    }
}

fn calculate_overlap(reader: BufReader<File>) -> (Vec<Overlap>, Vec<Overlap>) {
    reader
        .lines()
        .filter_map(|l| l.ok())
        .map(|line| {
            line.split(',')
                .map(|s| {
                    let r = s
                        .split('-')
                        .map(|i| i.parse::<i32>().expect("Assignments must be ints"))
                        .collect::<Vec<i32>>();
                    r[0]..=r[1]
                })
                .collect::<Vec<RangeInclusive<i32>>>()
        })
        .map(|r| r[0].overlaps(&r[1]))
        .filter(|o| !matches!(o, Overlap::None))
        .partition(|o| matches!(o, Overlap::Full))
}
